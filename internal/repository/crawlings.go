package repository

import "github.com/1000king/handover/internal/domain"

type CrawlThreadsRepository interface {
	FindLastIndex() (int, error)
	InsertThread(startIndex, lastIndex int) error
}

type CrawlKeywordsRepository interface {
	InsertKeyword(keyword string, registeredIndex int) error
	FindLiveKeywords() ([]*domain.CrawlKeyword, error)
}

type CrawlProductsRepository interface{}
