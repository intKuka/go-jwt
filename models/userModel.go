package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserId string             `bson:"user_id"`
	Session
}
