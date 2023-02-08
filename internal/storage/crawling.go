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
	"go.uber.org/zap"
)

type crawlThreadRepo struct {
	col *mongo.Collection
}

func (repo *crawlThreadRepo) InsertThread(newThread *domain.CrawlThread) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := repo.col.InsertOne(ctx, newThread)
	if err != nil {
		zap.Error(err)
		return err
	}

	return nil
}

func (repo *crawlThreadRepo) FindLastIndex() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	findOptions := options.FindOne().SetSort(bson.D{{Key: "lastIndex", Value: -1}})

	var result *domain.CrawlThread

	if err := repo.col.FindOne(ctx, bson.M{}, findOptions).Decode(&result); err != nil {
		return 0, nil
	}

	return result.LastIndex, nil
}

func MongoCrawlThreadsRepo(conn *MongoDB) repository.CrawlThreadsRepository {
	return &crawlThreadRepo{
		col: conn.crawlThreadCol,
	}
}

type crawlKeywordRepo struct {
	col *mongo.Collection
}

func (repo *crawlKeywordRepo) FindLiveKeywords() ([]*domain.CrawlKeyword, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var result []*domain.CrawlKeyword
	cursor, err := repo.col.Find(ctx, bson.M{"isalive": true})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *crawlKeywordRepo) InsertKeyword(keyword string, registeredIndex int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	now := time.Now()

	result := domain.CrawlKeyword{
		Keyword:         keyword,
		IsAlive:         true,
		RegisteredIndex: registeredIndex,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	_, err := repo.col.InsertOne(ctx, &result)
	if err != nil {
		return err
	}

	return nil
}

func MongoCrawlKeywordsRepo(conn *MongoDB) repository.CrawlKeywordsRepository {
	return &crawlKeywordRepo{
		col: conn.crawlKeywordCol,
	}
}

type crawlProductRepo struct {
	col *mongo.Collection
}

// List implements repository.CrawlProductsRepository
func (repo *crawlProductRepo) List(filter *domain.CrawlProductFilter, offset, limit int) ([]*domain.CrawlProduct, int, error) {
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

	var pds []*domain.CrawlProduct
	err = cursor.All(ctx, &pds)
	if err != nil {
		return nil, 0, err
	}
	return pds, int(totalCount), nil
}

// Insert implements repository.CrawlProductsRepository
func (repo *crawlProductRepo) Insert(pd *domain.CrawlProduct) (*domain.CrawlProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	result, err := repo.col.InsertOne(ctx, pd)
	if err != nil {
		return nil, err
	}

	pd.ID = result.InsertedID.(primitive.ObjectID)

	return pd, nil
}

// Update implements repository.CrawlProductsRepository
func (repo *crawlProductRepo) Update(pd *domain.CrawlProduct) (*domain.CrawlProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": pd.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &pd}, opts); err != nil {
		return nil, err
	}

	var updatedProduct *domain.CrawlProduct
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedProduct); err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func MongoCrawlProductsRepo(conn *MongoDB) repository.CrawlProductsRepository {
	return &crawlProductRepo{
		col: conn.crawlProductCol,
	}
}
