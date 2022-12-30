package system

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"server/common/logs"
	"server/common/model"
	"server/common/response"
)

func Reset(ctx context.Context, c *app.RequestContext) {
	err := model.SqliteCreateTable(&logs.Logs{})
	if err != nil {
		response.ErrorRes(c, "创建log表错误", err)
		return
	}
	response.SuccessRes(c, "重置成功", nil)
}
