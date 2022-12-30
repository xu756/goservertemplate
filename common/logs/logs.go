package logs

import (
	"gorm.io/datatypes"
	"time"
)

type Logs struct {
	Id        int            `json:"id" gorm:"primary_key,comment:'主键'"`
	Username  string         `json:"username" gorm:"comment:'访问用户名',default:'未知用户'"`
	Ip        string         `json:"ip" gorm:"comment:'访问ip'"`
	Code      int            `json:"code" gorm:"comment:'返回状态码'"`
	Type      string         `json:"type" gorm:"comment:'分类'"`
	Msg       string         `json:"msg" gorm:"comment:'返回信息'"`
	Data      datatypes.JSON `json:"data" gorm:"comment:'返回数据'"`
	Url       string         `json:"url" gorm:"comment:'访问url'"`
	CreatedAt time.Time      `json:"created_at" gorm:"comment:'创建时间'"`
}
