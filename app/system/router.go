package system

import (
	"github.com/gin-gonic/gin"
)

func SystemRouter(system *gin.RouterGroup) {
	system.GET("reset", Reset)
}
