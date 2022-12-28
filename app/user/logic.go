package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"server/common/response"
)

// 登录
func Login(ctx context.Context, c *app.RequestContext) {
	var login struct {
		Username string `json:"username,required" vd:"$!=''"`
		Password string `json:"password,required" vd:"$!=''"`
	}
	if err := c.BindAndValidate(&login); err != nil {
		log.Print(err)
		response.ErrorRes(c, "参数错误", nil)
		return
	}
	response.SuccessRes(c, "登录成功", login)
}
