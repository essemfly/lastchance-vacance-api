package crawler

import "github.com/1000king/handover/config"

const (
	chunkSize        = 1000
	numWorkers       = 5
	globalStartIndex = 1000001
)

func DanggnCrawler() {
	lastIndex := getLastIndex()
	for IsIndexExists(lastIndex + chunkSize) {
		startIndex := lastIndex + 1
		lastIndex = startIndex + chunkSize
		config.Repo.CrawlThreads.InsertThread(startIndex, lastIndex)
		crawlDanggnIndex(startIndex, lastIndex)
	}
}

func getLastIndex() int {
	lastIndex, err := config.Repo.CrawlThreads.FindLastIndex()
	if err != nil || lastIndex == 0 {
		return globalStartIndex
	}
	return lastIndex
}
