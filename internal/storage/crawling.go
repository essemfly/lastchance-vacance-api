package storage

import (
	"github.com/1000king/handover/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type crawlThreadRepo struct {
	col *mongo.Collection
}

func MongoCrawlThreadsRepo(conn *MongoDB) repository.CrawlThreadsRepository {
	return &productRepo{
		col: conn.crawlThreadCol,
	}
}

type crawlKeywordRepo struct {
	col *mongo.Collection
}

func MongoCrawlKeywordsRepo(conn *MongoDB) repository.CrawlKeywordsRepository {
	return &productRepo{
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
