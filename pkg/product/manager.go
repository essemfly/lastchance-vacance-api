package product

import (
	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
)

type AddProductRequest struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Images          []string `json:"images"`
	OriginalPrice   int      `json:"original_price"`
	DiscountedPrice int      `json:"discounted_price"`
	Outlink         string   `json:"outlink"`
}

func UpsertProductInCrawled(crawlPd *domain.CrawlProduct) (*domain.Product, error) {
	statusMatching := map[domain.DanggnStatus]domain.ProductStatus{
		domain.DANGGN_STATUS_SALE:    domain.PRODUCT_STATUS_SALE,
		domain.DANGGN_STATUS_SOLDOUT: domain.PRODUCT_STATUS_SOLDOUT,
		domain.DANGGN_STATUS_UNKNOWN: domain.PRODUCT_STATUS_UNKNOWN,
	}

	defaultImage := ""
	if len(crawlPd.Images) > 0 {
		defaultImage = crawlPd.Images[0]
	}

	newPd := &domain.Product{
		CrawlProductID:  crawlPd.ID,
		Name:            crawlPd.Name,
		UploadType:      domain.PRODUCT_TYPE_DANGGN,
		Description:     crawlPd.Description,
		Images:          crawlPd.Images,
		DefaultImage:    defaultImage,
		WrittenAddr:     crawlPd.SellerRegionName,
		Status:          statusMatching[crawlPd.Status],
		DiscountedPrice: crawlPd.Price,
		Outlink:         crawlPd.Url,
		WrittenAt:       crawlPd.WrittenAt,
		ViewCounts:      crawlPd.ViewCounts,
		LikeCounts:      crawlPd.LikeCounts,
		ChatCounts:      crawlPd.ChatCounts,
	}

	pd, err := config.Repo.Products.GetByCrawlID(crawlPd.ID)
	if err != nil {
		newPd, _ := config.Repo.Products.Insert(newPd)
		AddKeywordProduct(newPd)
		return newPd, nil

	}

	newPd.ID = pd.ID
	newPd.CreatedAt = pd.CreatedAt
	return config.Repo.Products.Update(newPd)

}

func AddProductManual(addPdReq *AddProductRequest) (*domain.Product, error) {
	return nil, nil
}
