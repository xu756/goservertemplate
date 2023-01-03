package logic

import (
	"context"
	"goservertemplate/common/errorx"

	"goservertemplate/app/user/api/internal/svc"
	"goservertemplate/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info() (resp *types.Info, err *errorx.CodeError) {
	return &types.Info{
		Username: l.ctx.Value("username").(string),
		Password: "123456",
	}, errorx.NewCodeError(200, "OK")
}
