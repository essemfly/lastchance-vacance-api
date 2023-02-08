package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductStatus string

const (
	PRODUCT_STATUS_SOLDOUT ProductStatus = "SOLDOUT"
	PRODUCT_STATUS_SALE    ProductStatus = "SALE"
	PRODUCT_STATUS_UNKNOWN ProductStatus = "UNKNOWN"
)

type ProductType string

const (
	PRODUCT_TYPE_DANGGN ProductType = "DANGGN"
	PRODUCT_TYPE_DIRECT ProductType = "DIRECT"
)

type ProductFilter struct {
	Status        ProductStatus
	SearchKeyword string
}

type Product struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CrawlProductID  primitive.ObjectID `json:"crawl_product_id"`
	Name            string             `json:"name"`
	UploadType      ProductType        `json:"upload_type"`
	Description     string             `json:"description"`
	Images          []string           `json:"images"`
	Status          ProductStatus      `json:"status"`
	OriginalPrice   int                `json:"original_price"`
	DiscountedPrice int                `json:"discounted_price"`
	Outlink         string             `json:"outlink"`
	ViewCounts      int                `json:"view_counts"` // 현재는 당근에서 가져오는 것
	LikeCounts      int                `json:"like_counts"` // 현재는 당근에서 가져오는 것
	ChatCounts      int                `json:"chat_counts"` // 현재는 당근에서 가져오는 것
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	DeletedAt       time.Time          `json:"deleted_at"`
}
