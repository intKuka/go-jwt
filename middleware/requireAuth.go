package middleware

import (
	"context"
	"fmt"
	"jwt-project/consts"
	"jwt-project/initializers"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func RequireAuth(c *gin.Context) {
	tokenString := c.GetHeader(consts.AuthHeader)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token is expired",
			})
		}

		coll := initializers.Client.Database(consts.DbName).Collection(consts.UsersCollection)

		if err := coll.FindOne(context.TODO(), bson.D{{Key: "user_id", Value: claims["sub"]}}).Err(); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user_id", claims["sub"])

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
