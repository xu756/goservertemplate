// Code generated by goctl. DO NOT EDIT.
// Source: user-rpc.proto

package server

import (
	"context"

	"goservertemplate/app/user/rpc/internal/logic"
	"goservertemplate/app/user/rpc/internal/svc"
	"goservertemplate/app/user/rpc/user"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) Login(ctx context.Context, in *user.User) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}