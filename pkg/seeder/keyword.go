package pkg

import (
	"log"
	"time"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func AddKeywordSeed(startIndex int) {
	keywords := []domain.CrawlKeyword{
		{
			ID:              primitive.NewObjectID(),
			Keyword:         "숙박",
			IsAlive:         true,
			RegisteredIndex: startIndex,
			LastIndex:       startIndex,
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
		},
		{
			ID:              primitive.NewObjectID(),
			Keyword:         "호텔",
			IsAlive:         true,
			RegisteredIndex: startIndex,
			LastIndex:       startIndex,
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
		},
		{
			ID:              primitive.NewObjectID(),
			Keyword:         "양도",
			IsAlive:         true,
			RegisteredIndex: startIndex,
			LastIndex:       startIndex,
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
		},
	}

	for _, keyword := range keywords {
		err := config.Repo.CrawlKeywords.InsertKeyword(keyword.Keyword, keyword.RegisteredIndex)
		if err != nil {
			// config.Logger.Error("failed to insert keyword", zap.Error(err))
			log.Fatalln("failed to insert keyword", zap.Error(err))
			return
		}
	}

}
