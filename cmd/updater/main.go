package main

import (
	"github.com/1000king/handover/cmd"
	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"go.uber.org/zap"
)

func main() {
	cmd.InitBase()

	filter := &domain.CrawlProductFilter{
		Status: domain.DANGGN_STATUS_SALE,
	}

	offset, limit := 0, 100
	pds, total, err := config.Repo.CrawlProducts.List(filter, offset, limit)
	if err != nil {
		config.Logger.Error("failed to list product", zap.Error(err))
	}

	for _, pd := range pds {
		updateProductStatus(pd)
	}

	if total > limit {
		for i := 1; i < total/limit; i++ {
			offset = i * limit
			pds, _, err = config.Repo.CrawlProducts.List(filter, offset, limit)
			if err != nil {
				config.Logger.Error("failed to list product", zap.Error(err))
			}
			for _, pd := range pds {
				updateProductStatus(pd)
			}
		}
	}
}

func updateProductStatus(pd *domain.CrawlProduct) {
	_, err := config.Repo.CrawlProducts.Update(pd)
	if err != nil {
		config.Logger.Error("failed to update product", zap.Error(err))
	}

	pds, err := config.Repo.Products.ListByCrawlID(pd.ID)
	if err != nil {
		config.Logger.Error("failed to list product by crawled id", zap.Error(err))
	}

	statusMatching := map[domain.DanggnStatus]domain.ProductStatus{
		domain.DANGGN_STATUS_SALE:    domain.PRODUCT_STATUS_SALE,
		domain.DANGGN_STATUS_SOLDOUT: domain.PRODUCT_STATUS_SOLDOUT,
		domain.DANGGN_STATUS_UNKNOWN: domain.PRODUCT_STATUS_UNKNOWN,
	}

	for _, p := range pds {
		p.Status = statusMatching[pd.Status]
		_, err := config.Repo.Products.Update(p)
		if err != nil {
			config.Logger.Error("failed to update product", zap.Error(err))
		}
	}
}
