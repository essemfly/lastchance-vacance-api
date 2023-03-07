package routes

import (
	"net/http"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListKeywords(c echo.Context) error {
	userIdStr := c.Get("user").(*jwt.Token)
	claims := userIdStr.Claims.(*JwtClaim)

	keywords, err := config.Repo.Keywords.List(claims.UserId)
	if err != nil {
		panic(err)
	}

	if len(keywords) < 1 {
		return c.JSON(http.StatusOK, []string{})
	}

	return c.JSON(http.StatusOK, keywords)
}

func InsertKeyword(c echo.Context) error {
	userIdStr := c.Get("user").(*jwt.Token)
	claims := userIdStr.Claims.(*JwtClaim)
	userId, _ := primitive.ObjectIDFromHex(claims.UserId)
	keyword := c.FormValue("keyword")

	keywordObj := &domain.Keyword{
		UserID:  userId,
		Keyword: keyword,
	}

	newOrder, err := config.Repo.Keywords.Insert(keywordObj)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, newOrder)
}

func UpdateKeyword(c echo.Context) error {
	keywordIDStr := c.FormValue("keywordid")
	keyword, err := config.Repo.Keywords.Get(keywordIDStr)
	if err != nil {
		panic(err)
	}

	keyword.IsLive = false

	updatedKeyword, err := config.Repo.Keywords.Update(keyword)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, updatedKeyword)
}

func ListKeywordProducts(c echo.Context) error {
	userIdStr := c.Get("user").(*jwt.Token)
	claims := userIdStr.Claims.(*JwtClaim)

	pds, err := config.Repo.KeywordProducts.List(claims.UserId)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, pds)
}
