package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Mobile     string             `json:"mobile"`
	Address    string             `json:"address"`
	DeviceUUID string             `json:"device_uuid"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

type UserLike struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"user_id"`
	ProductId primitive.ObjectID `json:"product_id"`
	IsLiked   bool               `json:"is_liked"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type UserLikeFilter struct {
	UserId primitive.ObjectID
}
