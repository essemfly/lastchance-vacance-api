package storage

import (
	"context"
	"time"

	"github.com/1000king/handover/internal/domain"
	"github.com/1000king/handover/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type orderRepo struct {
	col *mongo.Collection
}

func (repo *orderRepo) Insert(order *domain.Order) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	order.ID = primitive.NewObjectID()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	_, err := repo.col.InsertOne(ctx, order)
	if err != nil {
		zap.Error(err)
		return nil, err
	}

	return order, nil
}

func (repo *orderRepo) List(filter *domain.OrderFilter) ([]*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cursor, err := repo.col.Find(ctx, bson.M{"userid": filter.UserId})
	if err != nil {
		return nil, err
	}

	var orders []*domain.Order
	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func MongoOrdersRepo(conn *MongoDB) repository.OrdersRepository {
	return &orderRepo{
		col: conn.orderCol,
	}
}
