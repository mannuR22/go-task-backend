package dao

import (
	"context"

	mongo "github.com/mannuR22/go-task-backend.git/dao/connection"
	list "github.com/mannuR22/go-task-backend.git/models"
)

func createList(userList list.ListModel) error {
	client, err := mongo.Connect()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	// Get a handle to the "test" database
	db := client.Database("taskdb")

	// Get a handle to the "users" collection
	lists := db.Collection("lists")

	// Insert a new document
	var listDao := list.ListDAO
	_, err = lists.InsertOne(context.Background(), userList)
	if err != nil {
		return err
	}

	return nil
}

func readList() {

}

func updateList() {}

func deleteListItems() {}
