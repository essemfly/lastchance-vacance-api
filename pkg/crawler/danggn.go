package crawler

import (
	"github.com/1000king/handover/config"
	"go.uber.org/zap"
)

func IsIndexExists(index int) bool {
	return true
}

func crawlDanggnIndex(startIndex, lastIndex int) {
	_, err := config.Repo.CrawlKeywords.FindLiveKeywords()
	if err != nil {
		config.Logger.Error("failed to find live keywords", zap.Error(err))
	}

}

func crawlPage(index int) {

}
