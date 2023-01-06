package storage

import (
	"github.com/1000king/handover/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type notificationRepo struct {
	col *mongo.Collection
}

func MongoNotificationsRepo(conn *MongoDB) repository.NotificationsRepository {
	return &productRepo{
		col: conn.notificationCol,
	}
}
