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

func (repo *productRepo) Get(ID primitive.ObjectID) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	mongoFilter := bson.M{
		"_id": ID,
	}

	var pd *domain.Product

	if err := repo.col.FindOne(ctx, mongoFilter).Decode(&pd); err != nil {
		return nil, err
	}
	return pd, nil
}

func (repo *productRepo) ListByCrawlID(crawlID primitive.ObjectID) ([]*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	mongoFilter := bson.M{
		"craulid": crawlID,
	}

	cursor, err := repo.col.Find(ctx, mongoFilter)
	if err != nil {
		return nil, err
	}

	var pds []*domain.Product
	err = cursor.All(ctx, &pds)
	if err != nil {
		return nil, err
	}
	return pds, nil
}

func (repo *productRepo) List(filter *domain.ProductFilter, offset, limit int) ([]*domain.Product, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(bson.D{{Key: "createdat", Value: -1}})
	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))

	twoWeeksAgo := time.Now().AddDate(0, 0, -14)
	mongoFilter := bson.M{}

	if filter.SearchKeyword != "" {
		mongoFilter["$or"] = []bson.M{
			{"name": primitive.Regex{Pattern: filter.SearchKeyword, Options: "i"}},
			{"description": primitive.Regex{Pattern: filter.SearchKeyword, Options: "i"}},
			{"writtenaddr": primitive.Regex{Pattern: filter.SearchKeyword, Options: "i"}},
		}
	} else {
		mongoFilter["status"] = domain.PRODUCT_STATUS_SALE
		mongoFilter["writtenat"] = bson.M{"$gte": twoWeeksAgo}
	}

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

func (repo *productRepo) Insert(pd *domain.Product) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	pd.ID = primitive.NewObjectID()
	pd.CreatedAt = time.Now()
	pd.UpdatedAt = time.Now()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"crawlproductid": pd.CrawlProductID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": pd}, opts); err != nil {
		return nil, err
	}

	return pd, nil
}

func (repo *productRepo) Update(pd *domain.Product) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	pd.UpdatedAt = time.Now()
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

func (repo *productRepo) Remove(productID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"_id": productID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"deletedat": time.Now()}}); err != nil {
		return err
	}
	return nil
}

func MongoProductsRepo(conn *MongoDB) repository.ProductsRepository {
	return &productRepo{
		col: conn.productCol,
	}
}
