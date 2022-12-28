package user

import (
	"github.com/cloudwego/hertz/pkg/route"
)

func UserRouter(user *route.RouterGroup) {
	user.POST("login", Login)
}
