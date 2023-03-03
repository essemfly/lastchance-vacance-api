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
	Get(userID primitive.ObjectID, productID primitive.ObjectID) (*domain.UserLike, error)
	List(filter *domain.UserLikeFilter) ([]*domain.UserLike, error)
	Upsert(primitive.ObjectID, *domain.Product) (*domain.UserLike, error)
}
