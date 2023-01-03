package logic

import (
	"context"
	"goservertemplate/common/errorx"

	"goservertemplate/app/user/api/internal/svc"
	"goservertemplate/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfo, err *errorx.CodeError) {

	return &types.UserInfo{
		Username: l.ctx.Value("username").(string),
	}, errorx.NewCodeError(1001, "test")
}
