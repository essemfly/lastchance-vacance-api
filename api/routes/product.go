package routes

import (
	"net/http"
	"strconv"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetProductResponse struct {
	Product  *domain.Product  `json:"product"`
	UserLike *domain.UserLike `json:"userLike"`
}

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

	userIdStr := c.Get("user").(*jwt.Token)
	claims := userIdStr.Claims.(*JwtClaim)
	userID, _ := primitive.ObjectIDFromHex(claims.UserId)

	userLike, err := config.Repo.UserLikes.Get(userID, productID)
	if err != nil {
		panic(err)
	}

	productResponse := &GetProductResponse{
		Product:  pd,
		UserLike: userLike,
	}

	return c.JSON(http.StatusOK, productResponse)
}

func ListProducts(c echo.Context) error {
	offsetStr := c.QueryParam("page")
	limitStr := c.QueryParam("size")
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
