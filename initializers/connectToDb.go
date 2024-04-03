package initializers

import (
	"context"
	"fmt"
	"jwt-project/consts"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectToDb() {
	var err error
	uri := os.Getenv("MONGODB_URI")
	opts := options.Client().ApplyURI(uri)

	Client, err = mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	var result bson.M
	if err := Client.Database(consts.DbName).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
