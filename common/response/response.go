package response

import (
	"github.com/gin-gonic/gin"
	"server/common/logs"
	"server/common/method"
	"server/common/model"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Res(c *gin.Context, code int, msg string, data interface{}) *Response {
	go func() {
		model.GetSqlitedb().Create(&logs.Logs{
			Username: "admin",
			Ip:       c.ClientIP(),
			Code:     code,
			Type:     codes[code],
			Msg:      msg,
			Data:     method.GoBytes(data),
			Url:      c.Request.URL.Path,
		})

	}()
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func SuccessRes(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, Res(c, 200, msg, data))
}
func WarnRes(c *gin.Context, msg string, data interface{}) {
	c.JSON(300, Res(c, 400, msg, data))

}
func ErrorRes(c *gin.Context, msg string, data interface{}) {
	c.JSON(400, Res(c, 300, msg, data))
}
