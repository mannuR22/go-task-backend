package dao

import (
	"context"

	mongo "github.com/mannuR22/go-task-backend/dao/connection"
	"go.mongodb.org/mongo-driver/bson"
)

func createList() {
	client, err := mongo.Connect()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	// Get a handle to the "test" database
	db := client.Database("test")

	// Get a handle to the "users" collection
	users := db.Collection("users")

	// Insert a new document
	_, err = users.InsertOne(context.Background(), bson.M{"name": "Alice"})
	if err != nil {
		panic(err)
	}
}

func readList() {

}

func updateList() {}

func deleteListItems() {}
