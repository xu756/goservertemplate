package system

import (
	"github.com/gin-gonic/gin"
	"server/common/logs"
	"server/common/model"
	"server/common/response"
)

func Reset(c *gin.Context) {
	err := model.SqliteCreateTable(&logs.Logs{})
	if err != nil {
		response.ErrorRes(c, "创建log表错误", err)
		return
	}
	response.SuccessRes(c, "重置成功", nil)
}
