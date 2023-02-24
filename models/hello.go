package hello

type Hello struct {
	Message string `json:"message" bson:"message"`
}
