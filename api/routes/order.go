package routes

import (
	"net/http"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListOrders(c echo.Context) error {
	userIdStr := c.Get("user").(*jwt.Token)
	claims := userIdStr.Claims.(*JwtClaim)
	userId, _ := primitive.ObjectIDFromHex(claims.UserId)
	orderFilter := &domain.OrderFilter{
		UserId: userId,
	}

	orders, err := config.Repo.Orders.List(orderFilter)
	if err != nil {
		panic(err)
	}

	pds := []*domain.Product{}
	for _, order := range orders {
		pd, err := config.Repo.Products.Get(order.ProductId)
		if err != nil {
			panic(err)
		}
		pds = append(pds, pd)
	}

	return c.JSON(http.StatusOK, pds)
}

func CreateOrder(c echo.Context) error {
	userIdStr := c.Get("user").(*jwt.Token)
	claims := userIdStr.Claims.(*JwtClaim)
	userId, _ := primitive.ObjectIDFromHex(claims.UserId)
	pdIdStr := c.FormValue("productid")
	pdId, _ := primitive.ObjectIDFromHex(pdIdStr)
	mobile := c.FormValue("mobile")

	pd, _ := config.Repo.Products.Get(pdId)
	order := &domain.Order{
		Product:   pd,
		ProductId: pdId,
		UserId:    userId,
		Mobile:    mobile,
	}

	// TODO: 이 product가 order를 insert할 수 있는지 확인할 필요가 있음
	newOrder, err := config.Repo.Orders.Insert(order)
	if err != nil {
		panic(err)
	}
	config.SendOrderCreateMessage(mobile, userId.Hex(), pd.ID.Hex(), pd.Outlink)

	return c.JSON(http.StatusOK, newOrder)
}
