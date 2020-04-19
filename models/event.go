package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID      string             `json:"userId" bson:"userId"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	StartTime   time.Time          `json:"startTime" bson:"startTime"`
	EndTime     time.Time          `json:"endTime" bson:"endTime"`
	Notes       string             `json:"notes" bson:"notes"`
}
