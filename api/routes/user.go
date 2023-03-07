package routes

import (
	"net/http"
	"time"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtClaim struct {
	UserId   string `json:"userId"`
	DeviceId string `json:"deviceId"`
	Sub      string `json:"sub"`
	jwt.StandardClaims
}

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

	claims := &JwtClaim{
		Sub:      viper.GetString("API_TOKEN_SUB"),
		UserId:   newUser.ID.Hex(),
		DeviceId: newUser.DeviceUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, t)
}

func LikeProduct(c echo.Context) error {
	productIDStr := c.FormValue("productid")
	productID, _ := primitive.ObjectIDFromHex(productIDStr)

	userIdStr := c.Get("user").(*jwt.Token)
	claims := userIdStr.Claims.(*JwtClaim)
	userId, _ := primitive.ObjectIDFromHex(claims.UserId)
	pd, err := config.Repo.Products.Get(productID)
	if err != nil {
		panic(err)
	}

	userLike, err := config.Repo.UserLikes.Upsert(userId, pd)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, userLike)
}

func ListLikeProducts(c echo.Context) error {
	userIdStr := c.Get("user").(*jwt.Token)
	claims := userIdStr.Claims.(*JwtClaim)
	userId, _ := primitive.ObjectIDFromHex(claims.UserId)

	userLikeFilter := &domain.UserLikeFilter{
		UserId: userId,
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
