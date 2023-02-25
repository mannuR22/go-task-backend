package dao

import (
	"context"
	"fmt"

	mongo "github.com/mannuR22/go-task-backend.git/dao/connection"
	models "github.com/mannuR22/go-task-backend.git/models"
	utils "github.com/mannuR22/go-task-backend.git/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateList(userList models.ListModel) (string, error) {

	client, err := mongo.Connect()

	if err != nil {
		return "", err
	}

	defer client.Disconnect(context.Background())

	db := client.Database("taskdb")

	// Get a handle to the "users" collection
	lists := db.Collection("lists")
	// Insert a new document

	var listDao models.ListDAO

	ref := utils.GetRefId()
	listDao.Reference = ref
	listDao.Task = userList.Task
	listDao.Movies = userList.Movies
	listDao.Books = userList.Books

	_, err = lists.InsertOne(context.Background(), listDao)
	if err != nil {
		return "", err
	}

	return ref, nil
}

func ReadList(refId string) (models.ListModel, error) {
	client, err := mongo.Connect()
	var listDao models.ListDAO
	var userList models.ListModel
	if err != nil {
		return userList, err
	}

	defer client.Disconnect(context.Background())

	db := client.Database("taskdb")

	// Get a handle to the "users" collection
	lists := db.Collection("lists")
	// Insert a new document

	filter := bson.M{"reference": refId}

	err = lists.FindOne(context.Background(), filter).Decode(&listDao)

	if err != nil {
		return userList, err
	}
	userList.Task = listDao.Task
	userList.Movies = listDao.Movies
	userList.Books = listDao.Books

	return userList, nil
}

func UpdateListItems(refId string, userList models.ListModel) error {
	client, err := mongo.Connect()
	fmt.Println("hello")
	var listDao models.ListDAO
	if err != nil {
		return err
	}

	defer client.Disconnect(context.Background())

	db := client.Database("taskdb")

	// Get a handle to the "users" collection
	lists := db.Collection("lists")
	// Insert a new document

	filter := bson.M{"reference": refId}

	err = lists.FindOne(context.Background(), filter).Decode(&listDao)

	if err != nil {
		return err
	}

	for _, book := range userList.Books {
		if !utils.IsStringInArray(listDao.Books, book) {
			listDao.Books = append(listDao.Books, book)
		}
	}

	for _, movie := range userList.Movies {
		if !utils.IsStringInArray(listDao.Movies, movie) {
			listDao.Movies = append(listDao.Movies, movie)
		}
	}

	for _, task := range userList.Task {
		if !utils.IsStringInArray(listDao.Task, task) {
			listDao.Task = append(listDao.Task, task)
		}
	}

	update := bson.M{"$set": listDao}

	// Update document
	_, err2 := lists.UpdateOne(context.Background(), filter, update)
	if err2 != nil {

		return err2
	}

	return nil

}

func DeleteListItems(refId string, userList models.ListModel) error {
	client, err := mongo.Connect()

	var listDao models.ListDAO
	if err != nil {
		return err
	}

	defer client.Disconnect(context.Background())

	db := client.Database("taskdb")

	// Get a handle to the "users" collection
	lists := db.Collection("lists")
	// Insert a new document

	filter := bson.M{"reference": refId}

	err = lists.FindOne(context.Background(), filter).Decode(&listDao)

	if err != nil {
		return err
	}
	var newList models.ListModel
	for _, book := range listDao.Books {
		if !utils.IsStringInArray(userList.Books, book) {
			newList.Books = append(newList.Books, book)
		}
	}

	for _, movie := range listDao.Movies {
		if !utils.IsStringInArray(userList.Movies, movie) {
			newList.Movies = append(newList.Movies, movie)
		}
	}

	for _, task := range listDao.Task {
		if !utils.IsStringInArray(userList.Task, task) {
			newList.Task = append(newList.Task, task)
		}
	}

	update := bson.M{"$set": newList}

	// Update document
	_, err2 := lists.UpdateOne(context.Background(), filter, update)
	if err2 != nil {

		return err2
	}

	return nil

}
