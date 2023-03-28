package routes

import (
	"net/http"
	"strconv"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDanggnProduct(c echo.Context) error {
	productIDStr := c.Param("id")
	productID, _ := primitive.ObjectIDFromHex(productIDStr)
	pd, err := config.Repo.CrawlProducts.Get(productID)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, pd)
}

func ListDanggnProducts(c echo.Context) error {
	offsetStr := c.QueryParam("offset")
	limitStr := c.QueryParam("limit")
	keyword := c.QueryParam("keyword")
	offset, limit := 0, 100

	productFilter := &domain.CrawlProductFilter{
		Keyword: keyword,
		Status:  domain.DANGGN_STATUS_ALL,
	}

	if offsetStr != "" {
		offset, _ = strconv.Atoi(offsetStr)
	}
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	products, _, err := config.Repo.CrawlProducts.List(productFilter, offset, limit)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, products)
}
