package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/golang-jwt/jwt"
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

	claims := map[string]interface{}{
		"sub":      os.Getenv("API_TOKEN_SUB"),
		"userId":   newUser.ID,
		"deviceId": newUser.DeviceUUID,
	}

	accessToken, err := generateToken(c, claims)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, accessToken)
}

func LikeProduct(c echo.Context) error {
	productIDStr := c.FormValue("productid")
	productID, _ := primitive.ObjectIDFromHex(productIDStr)

	userIdStr, err := authHandler(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %s", err))
	}
	userId, _ := primitive.ObjectIDFromHex(userIdStr)
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
	userIdStr, err := authHandler(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %s", err))
	}
	userId, _ := primitive.ObjectIDFromHex(userIdStr)

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

// generateToken access 토큰과 refresh 토큰을 쌍으로 반환한다
func generateToken(c echo.Context, claimsMap map[string]interface{}) (string, error) {

	// access 토큰 생성: 유효기간 20분
	accessToken, err := createToken(
		c,
		claimsMap,
		time.Now().Add(time.Minute*20),
	)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// 단위 토큰 만들기
func createToken(c echo.Context, data map[string]interface{}, expire time.Time) (string, error) {

	// token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	for key, val := range data {
		claims[key] = val
	}
	claims["exp"] = expire.Unix()

	encToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return encToken, nil
}

func authHandler(c echo.Context) (string, error) {
	// Get JWT token from Authorization header
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return "", fmt.Errorf("Missing Authorization header")
	}

	// Parse JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Here you would typically implement logic to retrieve the key
		// for verifying the signature of the JWT token. This can be done
		// using a shared secret or by retrieving the public key of the
		// issuer from a well-known location.
		return []byte("my-secret-key"), nil
	})
	if err != nil {
		return "", fmt.Errorf("Invalid Authorization header")
	}

	// Extract claims from JWT token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("Invalid Authorization header")
	}

	// Access claims
	userID := claims["userId"].(string)

	return userID, nil
}
