package repository

import (
	"github.com/1000king/handover/internal/domain"
)

type KeywordsRepository interface {
	Get(keywordID string) (*domain.Keyword, error)
	Insert(*domain.Keyword) (*domain.Keyword, error)
	Update(*domain.Keyword) (*domain.Keyword, error)
	List(userID string) ([]*domain.Keyword, error)
	ListAll() ([]*domain.Keyword, error)
}

type KeywordProductsRepository interface {
	Insert(pd *domain.Product, userID string, keyword string) (*domain.Product, error)
	List(userID string) ([]*domain.KeywordProduct, error)
}
