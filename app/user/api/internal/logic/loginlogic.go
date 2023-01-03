package logic

import (
	"context"
	"goservertemplate/app/user/api/internal/svc"
	"goservertemplate/app/user/api/internal/types"
	"goservertemplate/app/user/rpc/user"
	"goservertemplate/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *LoginLogic) Login(req *types.Login) (resp *types.Loginres, err *errorx.CodeError) {
	res, err2 := l.svcCtx.UserRpc.Login(l.ctx, &user.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err2 != nil {
		return nil, nil
	}
	return &types.Loginres{
		Token: res.Token,
	}, nil
}
