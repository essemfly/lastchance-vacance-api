package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Keyword struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Keyword   string             `json:"keyword"`
	UserID    primitive.ObjectID `json:"user_id"`
	IsLive    bool               `json:"is_live"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
