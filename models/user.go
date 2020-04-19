package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	UserName  string             `json:"userName" bson:"userName"`
	Password  string             `json:"password" bson:"password"`
	Friends   string             `json:"friends" bson:"friends"`
}
