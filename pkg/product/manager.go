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

func AddProductInCrawled(pd *domain.CrawlProduct) (*domain.Product, error) {
	statusMatching := map[domain.DanggnStatus]domain.ProductStatus{
		domain.DANGGN_STATUS_SALE:    domain.PRODUCT_STATUS_SALE,
		domain.DANGGN_STATUS_SOLDOUT: domain.PRODUCT_STATUS_SOLDOUT,
		domain.DANGGN_STATUS_UNKNOWN: domain.PRODUCT_STATUS_UNKNOWN,
	}

	defaultImage := ""
	if len(pd.Images) > 0 {
		defaultImage = pd.Images[0]
	}

	newPd := &domain.Product{
		CrawlProductID:  pd.ID,
		Name:            pd.Name,
		UploadType:      domain.PRODUCT_TYPE_DANGGN,
		Description:     pd.Description,
		Images:          pd.Images,
		DefaultImage:    defaultImage,
		WrittenAddr:     pd.SellerRegionName,
		Status:          statusMatching[pd.Status],
		DiscountedPrice: pd.Price,
		Outlink:         pd.Url,
		WrittenAt:       pd.WrittenAt,
		ViewCounts:      pd.ViewCounts,
		LikeCounts:      pd.LikeCounts,
		ChatCounts:      pd.ChatCounts,
	}

	return config.Repo.Products.Insert(newPd)
}

func AddProductManual(addPdReq *AddProductRequest) (*domain.Product, error) {
	return nil, nil
}
