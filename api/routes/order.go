package routes

import (
	"fmt"
	"net/http"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListOrders(c echo.Context) error {
	userIdStr, err := authHandler(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %s", err))
	}
	userId, _ := primitive.ObjectIDFromHex(userIdStr)
	orderFilter := &domain.OrderFilter{
		UserId: userId,
	}

	orders, err := config.Repo.Orders.List(orderFilter)
	if err != nil {
		panic(err)
	}

	orderWithProducts := []*domain.OrderWithProduct{}
	for _, order := range orders {
		pd, err := config.Repo.Products.Get(order.ProductId)
		if err != nil {
			panic(err)
		}
		orderWithProducts = append(orderWithProducts, &domain.OrderWithProduct{
			Order:   *order,
			Product: *pd,
		})
	}

	return c.JSON(http.StatusOK, orderWithProducts)
}

func CreateOrder(c echo.Context) error {
	userIdStr, err := authHandler(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %s", err))
	}
	userId, _ := primitive.ObjectIDFromHex(userIdStr)
	pdIdStr := c.FormValue("productid")
	pdId, _ := primitive.ObjectIDFromHex(pdIdStr)
	mobile := c.FormValue("mobile")

	order := &domain.Order{
		ProductId: pdId,
		UserId:    userId,
		Mobile:    mobile,
	}

	// TODO: 이 product가 order를 insert할 수 있는지 확인할 필요가 있음
	newOrder, err := config.Repo.Orders.Insert(order)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, newOrder)
}
