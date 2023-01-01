package router

import (
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gin-gonic/gin"
	"log"
	"server/app/system"
	"server/app/user"
)

func InitRouter() {
	router := gin.Default()
	router.Group("api").GET("ping", func(c *gin.Context) {
		log.Print("ping")
		c.JSON(consts.StatusOK, utils.H{
			"message": "pong",
		})
	})
	user.UserRouter(router.Group("/api/user/"))
	system.SystemRouter(router.Group("/api/system/"))
	router.Run(":5986")
}
