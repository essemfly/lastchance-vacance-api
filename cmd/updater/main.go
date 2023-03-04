package main

import (
	"strconv"
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

	for {
		crawlProductFilter := &domain.CrawlProductFilter{
			Status: domain.DANGGN_STATUS_SALE,
		}
		offset, limit := 0, 1000
		crawlPds, total, err := config.Repo.CrawlProducts.List(crawlProductFilter, offset, limit)
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
			product.AddProductInCrawled(pd)
		}

		if total > limit {
			for i := 1; i < total/limit; i++ {
				offset = i * limit
				crawlPds, _, err = config.Repo.CrawlProducts.List(crawlProductFilter, offset, limit)
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
					product.AddProductInCrawled(pd)
				}
			}
		}
		time.Sleep(10 * time.Minute)
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
