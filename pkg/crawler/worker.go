package crawler

import (
	"github.com/1000king/handover/config"
	"go.uber.org/zap"
)

const (
	chunkSize        = 1000
	numWorkers       = 5
	globalStartIndex = 530520000 // 2023-02-07 09:35:00
)

func DanggnCrawler() {
	lastIndex := getLastIndex()
	for IsIndexExists(lastIndex + chunkSize) {
		startIndex := lastIndex + 1
		lastIndex = startIndex + chunkSize

		keywords, err := config.Repo.CrawlKeywords.FindLiveKeywords()
		if err != nil {
			config.Logger.Error("failed to find live keywords", zap.Error(err))
			return
		}
		crawlDanggnIndex(keywords, startIndex, lastIndex)
	}
}

func getLastIndex() int {
	lastIndex, err := config.Repo.CrawlThreads.FindLastIndex()
	if err != nil || lastIndex == 0 {
		return globalStartIndex
	}
	return lastIndex
}
