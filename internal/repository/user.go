package repository

import (
	"github.com/1000king/handover/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersRepository interface {
	Get(ID primitive.ObjectID) (*domain.User, error)
	Upsert(*domain.User) (*domain.User, error)
}

type UserLikesRepository interface {
	List(filter *domain.UserLikeFilter) ([]*domain.UserLike, error)
	Upsert(*domain.User, *domain.Product) (*domain.UserLike, error)
}
