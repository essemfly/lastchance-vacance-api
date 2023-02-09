package routes

import (
	"net/http"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(c echo.Context) error {
	deviceUUID := c.FormValue("deviceuuid")
	mobile := c.FormValue("mobile")

	user := &domain.User{
		DeviceUUID: deviceUUID,
		Mobile:     mobile,
	}

	newUser, err := config.Repo.Users.Upsert(user)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, newUser)
}

func LikeProduct(c echo.Context) error {
	productIDStr := c.FormValue("productid")
	productID, _ := primitive.ObjectIDFromHex(productIDStr)
	userIDStr := c.FormValue("userid")
	userID, _ := primitive.ObjectIDFromHex(userIDStr)

	user, err := config.Repo.Users.Get(userID)
	if err != nil {
		panic(err)
	}
	pd, err := config.Repo.Products.Get(productID)
	if err != nil {
		panic(err)
	}

	userLike, err := config.Repo.UserLikes.Upsert(user, pd)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, userLike)
}

func ListLikeProducts(c echo.Context) error {
	userIDStr := c.FormValue("userid")
	userID, _ := primitive.ObjectIDFromHex(userIDStr)

	userLikeFilter := &domain.UserLikeFilter{
		UserId: userID,
	}

	userLikes, err := config.Repo.UserLikes.List(userLikeFilter)
	if err != nil {
		panic(err)
	}

	var pds []*domain.Product
	for _, ul := range userLikes {
		pd, err := config.Repo.Products.Get(ul.ProductId)
		if err != nil {
			panic(err)
		}
		pds = append(pds, pd)
	}

	return c.JSON(http.StatusOK, pds)
}
