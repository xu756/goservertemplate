package router

import (
	"github.com/labstack/echo"
)

func InItRouter() {
	router := echo.New()
	r := router.Group("/api/v1")
	r.GET("/index", func(c echo.Context) error {
		return c.JSON(200, router.Routes())
	})
	router.Logger.Fatal(router.Start(":5986"))
}
