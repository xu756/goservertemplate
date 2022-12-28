package response

import "github.com/cloudwego/hertz/pkg/app"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Res(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func SuccessRes(c *app.RequestContext, msg string, data interface{}) {
	c.JSON(200, Res(200, msg, data))
}
func WarnRes(c *app.RequestContext, msg string, data interface{}) {
	c.JSON(300, Res(400, msg, data))

}
func ErrorRes(c *app.RequestContext, msg string, data interface{}) {
	c.JSON(400, Res(300, msg, data))
}
