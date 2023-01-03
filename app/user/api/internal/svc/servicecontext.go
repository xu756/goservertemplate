package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goservertemplate/app/user/api/internal/config"
	"goservertemplate/app/user/rpc/userservice"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
	}
}
