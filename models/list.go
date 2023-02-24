package models

type ListModel struct {
	Task   []string `json:"task" bson:"task"`
	Movies []string `json:"movies" bosn:"movies"`
	Books  []string `json:"books" bson:"books"`
}
