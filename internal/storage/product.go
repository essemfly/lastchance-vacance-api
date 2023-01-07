package storage

import (
	"context"
	"time"

	"github.com/1000king/handover/internal/domain"
	"github.com/1000king/handover/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepo struct {
	col *mongo.Collection
}

func (repo *productRepo) List(filter *domain.ProductFilter, offset, limit int) ([]*domain.Product, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(bson.D{{Key: "_id", Value: 1}})
	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))

	mongoFilter := bson.M{}

	totalCount, _ := repo.col.CountDocuments(ctx, mongoFilter)
	cursor, err := repo.col.Find(ctx, mongoFilter, options)
	if err != nil {
		return nil, 0, err
	}

	var pds []*domain.Product
	err = cursor.All(ctx, &pds)
	if err != nil {
		return nil, 0, err
	}
	return pds, int(totalCount), nil
}

// Insert implements repository.ProductsRepository
func (repo *productRepo) Insert(pd *domain.Product) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := repo.col.InsertOne(ctx, pd)
	if err != nil {
		return nil, err
	}

	pd.ID = result.InsertedID.(primitive.ObjectID)

	return pd, nil
}

func (repo *productRepo) Update(pd *domain.Product) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": pd.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &pd}, opts); err != nil {
		return nil, err
	}

	var updatedProduct *domain.Product
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedProduct); err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func MongoProductsRepo(conn *MongoDB) repository.ProductsRepository {
	return &productRepo{
		col: conn.productCol,
	}
}
