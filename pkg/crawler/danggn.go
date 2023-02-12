package crawler

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

const (
	ProductURL = "https://www.daangn.com/articles/"
)

func crawlDanggnIndex(worker chan bool, done chan bool, keywords []*domain.CrawlKeyword, startIndex, lastIndex int) {
	// config.Logger.Info("start crawling danggn index", zap.Int("startIndex", startIndex), zap.Int("lastIndex", lastIndex))
	log.Println("start crawling danggn index", zap.Int("startIndex", startIndex), zap.Int("lastIndex", lastIndex))
	numMatchedProducts := 0
	for i := startIndex; i <= lastIndex; i++ {
		newProduct, err := crawlPage(i)
		if err != nil {
			if err.Error() == "Not Found" {
				continue
			}

			// config.Logger.Error("failed to crawl page", zap.Error(err))
			log.Fatalln("failed to crawl page", zap.Error(err))
			continue
		}

		pds := addProductKeywords(newProduct, keywords)
		if len(pds) == 0 {
			continue
		}

		for _, pd := range pds {
			config.Repo.CrawlProducts.Insert(pd)
			numMatchedProducts += 1
		}
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
	}
	config.Repo.CrawlThreads.InsertThread(newThreadResult)
	<-worker
	done <- true
}

func crawlPage(index int) (*domain.CrawlProduct, error) {
	if index%500 == 0 {
		// config.Logger.Info("start crawling danggn page", zap.Int("index", index))
		log.Println("start crawling danggn page", zap.Int("index", index))
	}
	c := colly.NewCollector(
		colly.AllowedDomains("www.daangn.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	c.SetRequestTimeout(10 * time.Second)

	indexStr := strconv.Itoa(index)
	url := ProductURL + indexStr

	newProduct := domain.CrawlProduct{
		ID:          primitive.NewObjectID(),
		DanggnIndex: indexStr,
		Url:         url,
	}

	c.OnHTML("head", func(e *colly.HTMLElement) {
		// Find the value of the product:availability meta header
		availability := e.ChildAttr("meta[name='product:availability']", "content")
		if availability == "oos" {
			newProduct.Status = domain.DANGGN_STATUS_SOLDOUT
		} else if availability == "instock" {
			newProduct.Status = domain.DANGGN_STATUS_SALE
		} else {
			newProduct.Status = domain.DANGGN_STATUS_UNKNOWN
		}
	})

	c.OnHTML("#article-images", func(e *colly.HTMLElement) {
		e.ForEach("img", func(_ int, e *colly.HTMLElement) {
			imageUrl := e.Attr("data-lazy")
			newProduct.Images = append(newProduct.Images, imageUrl)
		})
	})

	c.OnHTML("#article-profile", func(e *colly.HTMLElement) {
		nickName := e.ChildText("#nickname")
		regionName := e.ChildText("#region-name")
		temperature := e.ChildText("#temperature-wrap dd")

		newProduct.SellerNickName = nickName
		newProduct.SellerRegionName = regionName
		newProduct.SellerTemperature = temperature
	})

	c.OnHTML("#article-description", func(e *colly.HTMLElement) {
		categoryAndWrittenDate := e.ChildText("#article-category")
		spliter := "∙"
		categoryParsers := strings.Split(categoryAndWrittenDate, spliter)
		if len(categoryParsers) > 1 {
			if strings.Contains(categoryParsers[1], "시간") {
				numHours, _ := extractIntFromString(categoryParsers[1])
				now := time.Now()
				newProduct.WrittenAt = now.Add(-1 * time.Duration(numHours) * time.Hour)
			} else if strings.Contains(categoryParsers[1], "일") {
				numDays, _ := extractIntFromString(categoryParsers[1])
				now := time.Now()
				newProduct.WrittenAt = now.Add(-1 * time.Duration(numDays) * 24 * time.Hour)
			}
		}

		title := e.ChildText("#article-title")
		price := e.ChildText("#article-price")
		description := e.ChildText("#article-detail")
		articleCounts := e.ChildText("#article-counts")

		newProduct.Name = title
		newProduct.Price = ParsePriceString(price)
		newProduct.Description = description

		likeCount, viewCount, chatCount := ParseViewCountsString(articleCounts)
		newProduct.CrawlCategory = categoryParsers[0]
		newProduct.LikeCounts = likeCount
		newProduct.ViewCounts = viewCount
		newProduct.ChatCounts = chatCount
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return &newProduct, nil
}

func addProductKeywords(product *domain.CrawlProduct, keywords []*domain.CrawlKeyword) []*domain.CrawlProduct {
	pds := []*domain.CrawlProduct{}
	for _, keyword := range keywords {
		if strings.Contains(product.Name, keyword.Keyword) {
			product.Keyword = keyword.Keyword
			pds = append(pds, product)
		}
	}
	return pds
}
