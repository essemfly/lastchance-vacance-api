package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/1000king/handover/cmd"
	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/1000king/handover/pkg/crawler"
	"github.com/1000king/handover/pkg/product"
	"go.uber.org/zap"
)

func main() {
	cmd.InitBase()

	log.Println("start updating danggn")

	crawlProductFilter := &domain.CrawlProductFilter{
		Title:        "",
		Keyword:      "",
		KeywordGroup: "handover",
		Status:       domain.DANGGN_STATUS_SALE,
	}
	offset, limit := 0, 1000
	_, total, err := config.Repo.CrawlProducts.List(crawlProductFilter, offset, limit)
	if err != nil {
		config.Logger.Error("failed to list product", zap.Error(err))
	}

	for total > offset {
		crawlPds, _, err := config.Repo.CrawlProducts.List(crawlProductFilter, offset, limit)
		if err != nil {
			config.Logger.Error("failed to list product", zap.Error(err))
		}
		for _, pd := range crawlPds {
			danggnIndex, _ := strconv.Atoi(pd.DanggnIndex)
			updatedCrawlProduct, err := crawler.CrawlPage(danggnIndex)
			if err != nil {
				zap.Error(err)
				continue
			}
			updateCrawledProduct(pd, updatedCrawlProduct)
			if pd.KeywordGroup == "handover" && screenCrawledProdduct(pd) {
				product.AddProductInCrawled(pd)
			}
		}
		offset += limit
	}
}

func updateCrawledProduct(pd *domain.CrawlProduct, newPd *domain.CrawlProduct) {
	pd.Status = newPd.Status
	pd.ViewCounts = newPd.ViewCounts
	pd.ChatCounts = newPd.ChatCounts
	pd.LikeCounts = newPd.LikeCounts
	pd.Price = newPd.Price
	pd.UpdatedAt = time.Now()

	_, err := config.Repo.CrawlProducts.Update(pd)
	if err != nil {
		config.Logger.Error("failed to update product", zap.Error(err))
	}
}

func screenCrawledProdduct(pd *domain.CrawlProduct) bool {
	if strings.Contains(pd.Name, "구매") {
		return false
	}
	if strings.Contains(pd.Name, "포켓몬") {
		return false
	}
	if strings.Contains(pd.Name, "닌텐도") {
		return false
	}
	if strings.Contains(pd.Name, "젤다") {
		return false
	}
	if strings.Contains(pd.Name, "레고") {
		return false
	}
	if strings.Contains(pd.Name, "서스펜션") {
		return false
	}
	if strings.Contains(pd.Name, "삽니다") {
		return false
	}
	if strings.Contains(pd.Name, "구합니다") {
		return false
	}
	if strings.Contains(pd.Name, "대리결제") {
		return false
	}
	if strings.Contains(pd.Name, "구해봅니다") {
		return false
	}
	if strings.Contains(pd.Name, "구해요") {
		return false
	}
	return true
}
