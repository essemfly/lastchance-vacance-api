package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DanggnStatus int

const (
	SALE DanggnStatus = iota
	CLOSE
	HIDE
)

type CrawlThread struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StartIndex         int                `json:"start_index"`
	LastIndex          int                `json:"last_index"`
	Keywords           []string           `json:"keywords"`
	NumCrawledProducts int                `json:"num_crawled_products"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
}

type CrawlKeyword struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Keyword         string             `json:"keyword"`
	IsAlive         bool               `json:"is_alive"`
	RegisteredIndex int                `json:"registered_index"`
	LastIndex       int                `json:"last_index"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}

type CrawlProductFilter struct {
	Keyword string
	Status  DanggnStatus
}

type CrawlProduct struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DanggnIndex       string             `json:"danggn_index"`
	Keyword           string             `json:"keyword"`
	Name              string             `json:"name"`
	Description       string             `json:"description"`
	Price             int                `json:"price"`
	Images            []string           `json:"images"`
	Status            DanggnStatus       `json:"status"`
	Url               string             `json:"url"`
	ViewCounts        int                `json:"view_counts"`
	LikeCounts        int                `json:"like_counts"`
	CrawlCategory     string             `json:"crawl_category"`
	SellerNickName    string             `json:"seller_nickname"`
	SellerRegionName  string             `json:"seller_region_name"`
	SellerTemperature string             `json:"seller_temperature"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}
