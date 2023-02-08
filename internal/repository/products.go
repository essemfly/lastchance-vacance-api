package repository

import (
	"github.com/1000king/handover/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsRepository interface {
	ListByCrawlID(crawlID primitive.ObjectID) ([]*domain.Product, error)
	Insert(*domain.Product) (*domain.Product, error)
	Update(*domain.Product) (*domain.Product, error)
	List(filter *domain.ProductFilter, offset, limit int) ([]*domain.Product, int, error)
}
