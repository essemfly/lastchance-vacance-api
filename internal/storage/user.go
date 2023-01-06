package storage

import (
	"github.com/1000king/handover/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	col *mongo.Collection
}

func MongoUsersRepo(conn *MongoDB) repository.UsersRepository {
	return &productRepo{
		col: conn.userCol,
	}
}

type userLikeRepo struct {
	col *mongo.Collection
}

func MongoUserLikesRepo(conn *MongoDB) repository.UserLikesRepository {
	return &productRepo{
		col: conn.userLikeCol,
	}
}
