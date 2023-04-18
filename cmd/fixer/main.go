package main

import (
	"log"

	"github.com/1000king/handover/cmd"
	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
)

func main() {
	cmd.InitBase()

	crawlFilter := &domain.CrawlProductFilter{
		Status: domain.DANGGN_STATUS_ALL,
	}

	offset, limit := 8000, 110000
	pds, _, err := config.Repo.CrawlProducts.List(crawlFilter, offset, limit)
	if err != nil {
		log.Println("Hoit", err)
	}

	for idx, pd := range pds {
		if idx%1000 == 0 {
		}
		if pd.Keyword == "숙박" || pd.Keyword == "리즈트" || pd.Keyword == "팬션" || pd.Keyword == "펜션" {
			pd.KeywordGroup = "handover"
			config.Repo.CrawlProducts.Update(pd)
		} else {
			pd.KeywordGroup = "rovers"
			config.Repo.CrawlProducts.Update(pd)
		}
	}
}
