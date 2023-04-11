package main

import (
	"log"
	"net/http"

	"github.com/1000king/handover/api/routes"
	"github.com/1000king/handover/cmd"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/spf13/viper"
)

func main() {
	cmd.InitBase()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	logger, _ := zap.NewProduction()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/user", routes.RegisterUser)

	e.GET("/danggn/:id", routes.GetDanggnProduct)
	e.GET("/danggns", routes.ListDanggnProducts)

	r := e.Group("/api")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &routes.JwtClaim{},
		SigningKey: []byte(viper.GetString("JWT_SECRET")),
	}))

	r.GET("/product/:id", routes.GetProduct)
	r.GET("/products", routes.ListProducts)

	r.GET("/user", routes.GetUserInfo)
	r.PUT("/user", routes.UpdateUserInfo)
	r.GET("/user/likes", routes.ListLikeProducts)
	r.POST("/user/like", routes.LikeProduct)

	r.GET("/user/orders", routes.ListOrders)
	r.POST("/user/order", routes.CreateOrder)

	r.GET("/user/keywords", routes.ListKeywords)
	r.POST("/user/keyword", routes.InsertKeyword)
	r.PUT("/user/keyword", routes.UpdateKeyword)
	r.GET("/user/keyword/products", routes.ListKeywordProducts)

	port := viper.GetString("API_PORT")
	log.Println("PORT", port)

	e.Logger.Fatal(e.Start(":" + port))
}
