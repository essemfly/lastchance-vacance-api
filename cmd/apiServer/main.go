package main

import (
	"log"
	"net/http"

	"github.com/1000king/handover/api/routes"
	"github.com/1000king/handover/cmd"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {
	cmd.InitBase()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/danggn/:id", routes.GetDanggnProduct)
	e.GET("/danggns", routes.ListDanggnProducts)

	e.GET("/product/:id", routes.GetProduct)
	e.GET("/products", routes.ListProducts)

	e.POST("/user", routes.RegisterUser)
	e.POST("/user/likes", routes.ListLikeProducts)
	e.PUT("/user/like", routes.LikeProduct)

	e.POST("/orders", routes.ListOrders)
	e.POST("/order", routes.CreateOrder)

	port := viper.GetString("API_PORT")
	log.Println("PORT", port)
	e.Logger.Fatal(e.Start(":" + port))
}
