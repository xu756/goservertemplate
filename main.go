package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	r := router.Group("/api/v1/")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":5986")
}
