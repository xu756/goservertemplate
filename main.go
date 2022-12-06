package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	r := e.Group("/api/v1/")
	// Routes
	r.GET("", hello).Name = "hello"

	// Start server
	s := &http.Server{
		Addr:         ":8888",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}

// Handler
func hello(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
