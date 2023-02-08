package repository

import (
	"github.com/1000king/handover/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CrawlThreadsRepository interface {
	FindLastIndex() (int, error)
	InsertThread(*domain.CrawlThread) error
}

type CrawlKeywordsRepository interface {
	InsertKeyword(keyword string, registeredIndex int) error
	FindLiveKeywords() ([]*domain.CrawlKeyword, error)
}

type CrawlProductsRepository interface {
	Get(ID primitive.ObjectID) (*domain.CrawlProduct, error)
	Insert(*domain.CrawlProduct) (*domain.CrawlProduct, error)
	Update(*domain.CrawlProduct) (*domain.CrawlProduct, error)
	List(filter *domain.CrawlProductFilter, offset, limit int) ([]*domain.CrawlProduct, int, error)
}
