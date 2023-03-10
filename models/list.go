package models

type ListModel struct {
	Task   []string `json:"task" bson:"task"`
	Movies []string `json:"movies" bosn:"movies"`
	Books  []string `json:"books" bson:"books"`
}

type ListDAO struct {
	Reference string   `json:"reference" bson:"reference"`
	Task      []string `json:"task" bson:"task"`
	Movies    []string `json:"movies" bosn:"movies"`
	Books     []string `json:"books" bson:"books"`
}

type ListCreatedVM struct {
	RefId string `json:"refId" bson:"refId"`
}
