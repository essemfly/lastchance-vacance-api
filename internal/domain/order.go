package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus int

const (
	OrderStatusSoldout OrderStatus = iota
	OrderStatusPending
	OrderStatusSucceeded
)

type Order struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductId   string             `json:"product_id"`
	User        User               `json:"user"`
	Mobile      string             `json:"mobile"`
	OrderStatus OrderStatus        `json:"order_status"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
