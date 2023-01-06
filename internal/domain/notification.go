package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}
