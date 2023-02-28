package routes

import (
	"net/http"
	"strconv"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListProductsResponse struct {
	Products []*domain.Product `json:"products"`
	TotalCnt int               `json:"totalCnt"`
}

func GetProduct(c echo.Context) error {
	productIDStr := c.Param("id")
	productID, _ := primitive.ObjectIDFromHex(productIDStr)
	pd, err := config.Repo.Products.Get(productID)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, pd)
}

func ListProducts(c echo.Context) error {
	pageStr := c.QueryParam("page")
	sizeStr := c.QueryParam("size")
	search := c.QueryParam("search")
	page, size := 0, 1000

	productFilter := &domain.ProductFilter{
		SearchKeyword: search,
	}

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	if sizeStr != "" {
		size, _ = strconv.Atoi(sizeStr)
	}

	offset := page * size
	limit := size

	products, totalCnt, err := config.Repo.Products.List(productFilter, offset, limit)
	if err != nil {
		panic(err)
	}

	productsResponse := ListProductsResponse{
		TotalCnt: totalCnt,
		Products: products,
	}

	return c.JSON(http.StatusOK, productsResponse)
}
