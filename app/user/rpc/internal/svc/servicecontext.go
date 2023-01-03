package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"goservertemplate/app/user/model"
	"goservertemplate/app/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
