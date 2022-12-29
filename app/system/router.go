package system

import (
	"github.com/cloudwego/hertz/pkg/route"
)

func SystemRouter(system *route.RouterGroup) {
	system.GET("reset", Reset)
}
