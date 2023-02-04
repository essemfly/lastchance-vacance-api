package crawler

import (
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
	return true
}

func crawlDanggnIndex(startIndex, lastIndex int) {
	_, err := config.Repo.CrawlKeywords.FindLiveKeywords()
	if err != nil {
		config.Logger.Error("failed to find live keywords", zap.Error(err))
	}

}

func crawlPage(index int) {
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
		config.Logger.Error("error occurred in crawl afound ", zap.Error(err))
	}
}
