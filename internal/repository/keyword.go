package repository

import (
	"github.com/1000king/handover/internal/domain"
)

type KeywordsRepository interface {
	Insert(*domain.Keyword) (*domain.Keyword, error)
	Update(*domain.Keyword) (*domain.Keyword, error)
	List(userID string) ([]*domain.Keyword, error)
}
