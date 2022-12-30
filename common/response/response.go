package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"server/common/logs"
	"server/common/method"
	"server/common/model"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var codes = map[int]string{
	200: "成功",
	300: "警告",
	400: "错误",
}

func Res(c *app.RequestContext, code int, msg string, data interface{}) *Response {
	go func() {
		model.GetSqlitedb().Create(&logs.Logs{
			Username: string(c.Request.Host()),
			Ip:       c.ClientIP(),
			Code:     code,
			Type:     codes[code],
			Msg:      msg,
			Data:     method.GoBytes(data),
			Url:      string(c.Request.URI().PathOriginal()),
		})

	}()
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func SuccessRes(c *app.RequestContext, msg string, data interface{}) {
	c.JSON(200, Res(c, 200, msg, data))
}
func WarnRes(c *app.RequestContext, msg string, data interface{}) {
	c.JSON(300, Res(c, 400, msg, data))

}
func ErrorRes(c *app.RequestContext, msg string, data interface{}) {
	c.JSON(400, Res(c, 300, msg, data))
}
