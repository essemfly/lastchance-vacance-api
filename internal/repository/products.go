package repository

import (
	"github.com/1000king/handover/internal/domain"
)

type ProductsRepository interface {
	Insert(*domain.Product) (*domain.Product, error)
	Update(*domain.Product) (*domain.Product, error)
	List(filter *domain.ProductFilter, offset, limit int) ([]*domain.Product, int, error)
}
