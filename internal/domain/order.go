package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus string

const (
	ORDER_STATUS_FAILED   OrderStatus = "FAILED"
	ORDER_STATUS_PROGRESS OrderStatus = "PROGRESS"
	ORDER_STATUS_READY    OrderStatus = "READY"
	ORDER_STATUS_SUCCESS  OrderStatus = "SUCCESS"
)

type OrderFilter struct {
	UserId primitive.ObjectID
}

type Order struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductId   primitive.ObjectID `json:"product_id"`
	UserId      primitive.ObjectID `json:"user"`
	Mobile      string             `json:"mobile"`
	OrderStatus OrderStatus        `json:"order_status"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type OrderWithProduct struct {
	Order   Order
	Product Product
}
