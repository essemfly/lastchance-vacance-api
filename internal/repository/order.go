package repository

import "github.com/1000king/handover/internal/domain"

type OrdersRepository interface {
	Insert(*domain.Order) (*domain.Order, error)
	List(filter *domain.OrderFilter) ([]*domain.Order, error)
}
