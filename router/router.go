package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"server/app/user"
)

func InitRouter() {
	h := server.Default(server.WithHostPorts("0.0.0.0:5986"))
	h.Group("api").GET("ping", func(ctx context.Context, c *app.RequestContext) {
		log.Print("ping")
		c.JSON(consts.StatusOK, utils.H{
			"message": "pong",
		})
	})
	user.UserRouter(h.Group("/api/user/"))
	h.Spin()
}
