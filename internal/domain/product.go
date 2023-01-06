package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductStatus int

const (
	ProductStatusSoldout ProductStatus = iota
	ProductStatusInSale
	ProductStatusReservated
)

type Product struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	Images          []string           `json:"images"`
	Status          ProductStatus      `json:"status"`
	OriginalPrice   int                `json:"original_price"`
	DiscountedPrice int                `json:"discounted_price"`
	Outlink         string             `json:"outlink"`
	ViewCounts      int                `json:"view_counts"` // 현재는 당근에서 가져오는 것
	LikeCounts      int                `json:"like_counts"` // 현재는 당근에서 가져오는 것
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`

	// Seller
	// Category
	// Place
	// ViewCounts + LikeCounts => 따로 collections 만들 예정
}
