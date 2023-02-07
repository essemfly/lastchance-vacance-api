package crawler

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

const (
	ProductURL = "https://www.daangn.com/articles/"
)

func IsIndexExists(index int) bool {
	_, err := crawlPage(index)
	// Check two times, since the first time may fail due to the server's response.
	if err != nil {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(11)
		_, err = crawlPage(index + n)
		if err != nil {
			return false
		}
	}

	return true
}

func crawlDanggnIndex(keywords []*domain.CrawlKeyword, startIndex, lastIndex int) {
	numMatchedProducts := 0
	for i := startIndex; i <= lastIndex; i++ {
		newProduct, err := crawlPage(i)
		if err != nil {
			config.Logger.Error("failed to crawl page", zap.Error(err))
			continue
		}

		if !basicClassifier(newProduct, keywords) {
			continue
		}

		config.Repo.CrawlProducts.Insert(newProduct)
		numMatchedProducts += 1
	}

	keywordsStr := make([]string, 0)
	for _, keyword := range keywords {
		keywordsStr = append(keywordsStr, keyword.Keyword)
	}

	newThreadResult := &domain.CrawlThread{
		StartIndex:         startIndex,
		LastIndex:          lastIndex,
		Keywords:           keywordsStr,
		NumCrawledProducts: numMatchedProducts,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	config.Repo.CrawlThreads.InsertThread(newThreadResult)
}

func crawlPage(index int) (*domain.CrawlProduct, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.daangn.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	c.SetRequestTimeout(10 * time.Second)

	indexStr := strconv.Itoa(index)
	url := ProductURL + indexStr

	newProduct := domain.CrawlProduct{
		DanggnIndex: indexStr,
		Url:         url,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	c.OnHTML("#article-images", func(e *colly.HTMLElement) {
		e.ForEach("img", func(_ int, e *colly.HTMLElement) {
			imageUrl := e.Attr("src")
			newProduct.Images = append(newProduct.Images, imageUrl)
		})
	})

	c.OnHTML("#article-profile", func(e *colly.HTMLElement) {
		nickName := e.ChildText("nickname")
		regionName := e.ChildText("region-name")
		temperature := e.ChildText("#temperature-wrap dd")

		newProduct.SellerNickName = nickName
		newProduct.SellerRegionName = regionName
		newProduct.SellerTemperature = temperature
	})

	c.OnHTML("#article-description", func(e *colly.HTMLElement) {
		category := e.ChildText("#article-category")
		title := e.ChildText("#article-title")
		price := e.ChildText("#article-price")
		description := e.ChildText("#article-detail")
		articleCounts := e.ChildText("#article-counts")

		newProduct.Name = title
		newProduct.Price = ParsePriceString(price)
		newProduct.Description = description

		likeCount, viewCount := ParseViewCountsString(articleCounts)
		newProduct.CrawlCategory = category
		newProduct.LikeCounts = likeCount
		newProduct.ViewCounts = viewCount
	})

	err := c.Visit(url)
	if err != nil {
		config.Logger.Error("error occurred in crawlling ", zap.Error(err))
		return nil, err
	}

	return &newProduct, nil
}

func basicClassifier(product *domain.CrawlProduct, keywords []*domain.CrawlKeyword) bool {
	return true
}
