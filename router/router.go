package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
)

func InitRouter() {
	h := server.Default(server.WithHostPorts("0.0.0.0:5986"))
	router := h.Group("/api/")
	router.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		log.Print("ping")
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "pong",
		})
	})
	h.Spin()
}
