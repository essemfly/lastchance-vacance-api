package product

import (
	"log"
	"strings"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
)

func AddKeywordProduct(pd *domain.Product) {
	liveKeywords, err := config.Repo.Keywords.ListAll()
	if err != nil {
		log.Println("failed to list keywords", err)
		return
	}

	for _, keyword := range liveKeywords {
		if strings.Contains(pd.Name, keyword.Keyword) {
			config.Repo.KeywordProducts.Insert(pd, keyword.UserID.Hex(), keyword.Keyword)
		}
	}
}
