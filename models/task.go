package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID      string             `json:"userId" bson:"userId"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	TaskDate    time.Time          `json:"taskDate" bson:"taskDate"`
	Completed   bool               `json:"completed" bson:"completed"`
}
