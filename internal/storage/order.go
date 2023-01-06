package storage

import (
	"github.com/1000king/handover/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type crawlOrderRepo struct {
	col *mongo.Collection
}

func MongoOrdersRepo(conn *MongoDB) repository.OrdersRepository {
	return &productRepo{
		col: conn.orderCol,
	}
}
