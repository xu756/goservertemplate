package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InItRouter() {
	router := gin.Default()
	r := router.Group("/api/v1/")
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(200, r)
		log.Print(router.Routes())
	})
	router.Run(":5986")
}
