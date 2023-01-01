package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/common/response"
)

// 登录
func Login(c *gin.Context) {
	var login struct {
		Username string `json:"username,required" vd:"$!=''"`
		Password string `json:"password,required" vd:"$!=''"`
	}
	if err := c.Bind(&login); err != nil {
		log.Print(err)
		response.ErrorRes(c, "参数错误", nil)
		return
	}
	response.SuccessRes(c, "登录成功", login)
}
