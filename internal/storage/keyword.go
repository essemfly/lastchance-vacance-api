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

type keywordRepo struct {
	col *mongo.Collection
}

// List implements repository.KeywordsRepository
func (repo *keywordRepo) List(userID string) ([]*domain.Keyword, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	userObjId, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.M{
		"userid": userObjId,
	}

	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var keywords []*domain.Keyword
	err = cursor.All(ctx, &keywords)
	if err != nil {
		return nil, err
	}
	return keywords, nil
}

// Upsert implements repository.KeywordsRepository
func (repo *keywordRepo) Insert(keyword *domain.Keyword) (*domain.Keyword, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	keyword.ID = primitive.NewObjectID()
	keyword.CreatedAt = time.Now()
	keyword.UpdatedAt = time.Now()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": keyword.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": keyword}, opts); err != nil {
		return nil, err
	}

	return keyword, nil
}

func (repo *keywordRepo) Update(keyword *domain.Keyword) (*domain.Keyword, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	keyword.UpdatedAt = time.Now()
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": keyword.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &keyword}, opts); err != nil {
		return nil, err
	}

	var updatedKeyword *domain.Keyword
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedKeyword); err != nil {
		return nil, err
	}

	return updatedKeyword, nil
}

func MongoKeywordRepo(conn *MongoDB) repository.KeywordsRepository {
	return &keywordRepo{
		col: conn.keywordCol,
	}
}
