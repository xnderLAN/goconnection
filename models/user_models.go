package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Phone    string             `bson:"phone"`
	Email    string             `bson:"email,omitempty"`
	Tags     []string           `bson:"tags,omitempty"`
}

type Login struct{
	Email 	  		string	`json:"email"`
	Password_hash   string	`json:"`
	Remember		bool
}