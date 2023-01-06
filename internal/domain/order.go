package domain

import "time"

type OrderStatus int

const (
	OrderStatusSoldout OrderStatus = iota
	OrderStatusPending
	OrderStatusSucceeded
)

type Order struct {
	ID          string      `json:"id"`
	ProductId   string      `json:"product_id"`
	User        User        `json:"user"`
	Mobile      string      `json:"mobile"`
	OrderStatus OrderStatus `json:"order_status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
