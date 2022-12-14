package router

import (
	"github.com/gin-gonic/gin"
)

func InItRouter() {
	router := gin.Default()
	router.Use(JWT())
	r := router.Group("/api/v1")
	r.GET("/", func(c *gin.Context) {
		m, _ := c.Get("jwt")
		c.Header("token", m.(string))
		M, err := ParseToken(m.(string))
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "成功",
			"data": M,
		})
	})
	router.Run(":5986")
}
