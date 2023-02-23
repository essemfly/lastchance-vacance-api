package routes

import (
	"net/http"
	"strconv"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	offsetStr := c.QueryParam("offset")
	limitStr := c.QueryParam("limit")
	search := c.QueryParam("search")
	offset, limit := 0, 1000

	productFilter := &domain.ProductFilter{
		SearchKeyword: search,
	}

	if offsetStr != "" {
		offset, _ = strconv.Atoi(offsetStr)
	}
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	products, _, err := config.Repo.Products.List(productFilter, offset, limit)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, products)
}
