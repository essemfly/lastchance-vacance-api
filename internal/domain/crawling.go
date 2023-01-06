package domain

import "time"

type DanggnStatus int

const (
	SALE DanggnStatus = iota
	CLOSE
	HIDE
)

type CrawlThread struct {
	ID           string    `json:"id"`
	StartIndex   int       `json:"start_index"`
	CurrentIndex int       `json:"current_index"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CrawlKeyword struct {
	ID         string    `json:"id"`
	Keyword    string    `json:"keyword"`
	IsAlive    bool      `json:"is_alive"`
	StartIndex int       `json:"start_index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CrawlProduct struct {
	ID          string       `json:"id"`
	DanggnIndex string       `json:"danggn_index"`
	Keyword     string       `json:"keyword"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       int          `json:"price"`
	Images      []string     `json:"images"`
	Status      DanggnStatus `json:"status"`
	Url         string       `json:"url"`
	ViewCounts  int          `json:"view_counts"` // 현재는 당근에서 가져오는 것
	LikeCounts  int          `json:"like_counts"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
