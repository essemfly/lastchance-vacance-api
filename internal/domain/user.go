package domain

import "time"

type User struct {
	ID         string    `json:"id"`
	DeviceUUID string    `json:"device_uuid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserLikes struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	ProductId string    `json:"product_id"`
	IsLiked   bool      `json:"is_liked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
