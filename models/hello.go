package models

type Hello struct {
	Message string `json:"message" bson:"message"`
}
