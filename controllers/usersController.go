package controllers

import (
	"context"
	"jwt-project/consts"
	"jwt-project/initializers"
	"jwt-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(c *gin.Context) {
	coll := initializers.Client.Database(consts.DbName).Collection(consts.UsersCollection)
	doc := models.User{
		ID:     primitive.NewObjectID(),
		UserId: uuid.New().String(),
	}

	_, err := coll.InsertOne(context.TODO(), doc)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id": doc.UserId,
	})
}

func GetUsers(c *gin.Context) {
	coll := initializers.Client.Database(consts.DbName).Collection(consts.UsersCollection)

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
