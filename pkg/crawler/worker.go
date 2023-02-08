package crawler

import (
	"log"
	"sync"

	"github.com/1000king/handover/config"
	"go.uber.org/zap"
)

const (
	chunkSize        = 1000
	numWorkers       = 5
	GlobalStartIndex = 531250000 // 2023-02-08 15:00:00
)

func DanggnCrawler() {
	var wg sync.WaitGroup

	lastIndex := getLastIndex()
	for IsIndexExists(lastIndex + chunkSize) {
		startIndex := lastIndex + 1
		lastIndex = startIndex + chunkSize - 1

		keywords, err := config.Repo.CrawlKeywords.FindLiveKeywords()
		if err != nil {
			// config.Logger.Error("failed to find live keywords", zap.Error(err))
			log.Fatalln("failed to find live keywords", zap.Error(err))
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			crawlDanggnIndex(keywords, startIndex, lastIndex)
		}()
	}
	wg.Wait()
}

func getLastIndex() int {
	lastIndex, err := config.Repo.CrawlThreads.FindLastIndex()
	if err != nil || lastIndex == 0 {
		return GlobalStartIndex
	}
	return lastIndex
}
