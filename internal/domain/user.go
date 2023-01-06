package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DeviceUUID string             `json:"device_uuid"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

type UserLikes struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId    string             `json:"user_id"`
	ProductId string             `json:"product_id"`
	IsLiked   bool               `json:"is_liked"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
