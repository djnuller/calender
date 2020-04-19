package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pending struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID         string             `json:"userId" bson:"userId"`
	PendingFriends []string           `json:"pendingFriends" bson:"pendingFriends"`
	PendingEvents  []string           `json:"pendingEvents" bson:"pendingEvents"`
}
