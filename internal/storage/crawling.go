package storage

import (
	"context"
	"time"

	"github.com/1000king/handover/internal/domain"
	"github.com/1000king/handover/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type crawlThreadRepo struct {
	col *mongo.Collection
}

func (repo *crawlThreadRepo) InsertThread(startIndex int, lastIndex int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	now := time.Now()

	result := domain.CrawlThread{
		StartIndex: startIndex,
		LastIndex:  lastIndex,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	_, err := repo.col.InsertOne(ctx, &result)
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
	cursor, err := repo.col.Find(ctx, bson.M{"is_alive": true})
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

func MongoCrawlProductsRepo(conn *MongoDB) repository.CrawlProductsRepository {
	return &productRepo{
		col: conn.crawlProductCol,
	}
}
