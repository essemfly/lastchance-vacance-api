package main

import (
	"github.com/1000king/handover/cmd"
	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/1000king/handover/pkg/product"
	"go.uber.org/zap"
)

func main() {
	cmd.InitBase()

	crawlProductFilter := &domain.CrawlProductFilter{
		Status: domain.DANGGN_STATUS_SALE,
	}
	offset, limit := 0, 1000
	crawlPds, total, err := config.Repo.CrawlProducts.List(crawlProductFilter, offset, limit)
	if err != nil {
		config.Logger.Error("failed to list product", zap.Error(err))
	}
	for _, pd := range crawlPds {
		updateCrawledProductStatus(pd)
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
				updateCrawledProductStatus(pd)
				product.AddProductInCrawled(pd)
			}
		}
	}
}

func updateCrawledProductStatus(pd *domain.CrawlProduct) {
	_, err := config.Repo.CrawlProducts.Update(pd)
	if err != nil {
		config.Logger.Error("failed to update product", zap.Error(err))
	}
}
