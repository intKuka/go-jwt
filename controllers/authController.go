package controllers

import (
	"context"
	"jwt-project/consts"
	"jwt-project/initializers"
	"jwt-project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func GetToken(c *gin.Context) {
	var token models.Token
	var user models.User

	id := c.Query("user_id")
	coll := initializers.Client.Database(consts.DbName).Collection(consts.UsersCollection)

	if err := coll.FindOne(context.TODO(), bson.D{{Key: "user_id", Value: id}}).Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := token.CreateTokenPair(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	user.ExpiresAt = token.CreatedAt.Add(consts.RefreshTokenTTL)
	user.RefreshToken = string(token.RefreshTokenHash)

	coll.FindOneAndReplace(context.TODO(), bson.D{{Key: "_id", Value: user.ID}}, user)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  token.AccessToken,
		"refreshToken": token.RefreshToken,
	})
}

func RefreshToken(c *gin.Context) {
	var token models.Token
	var user models.User
	var body struct {
		RefreshToken string `json:"refresh_token"`
		UserId       string `json:"user_id"`
	}

	coll := initializers.Client.Database(consts.DbName).Collection(consts.UsersCollection)

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := coll.FindOne(context.TODO(), bson.D{{Key: "user_id", Value: body.UserId}}).Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.RefreshToken), []byte(body.RefreshToken)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid refresh token",
		})
		return
	}

	if time.Now().Unix() > user.ExpiresAt.Unix() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Token is expired",
		})
		return
	}

	if err := token.CreateTokenPair(c, body.UserId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.ExpiresAt = token.CreatedAt.Add(consts.RefreshTokenTTL)
	user.RefreshToken = string(token.RefreshTokenHash)

	coll.FindOneAndReplace(context.TODO(), bson.D{{Key: "_id", Value: user.ID}}, user)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  token.AccessToken,
		"refreshToken": token.RefreshToken,
	})
}

func ValidateAuth(c *gin.Context) {
	id, _ := c.Get("user_id")

	c.JSON(http.StatusOK, gin.H{
		"message": "user " + id.(string) + " logged in",
	})
}
