package user

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(user *gin.RouterGroup) {
	user.POST("login", Login)
}
