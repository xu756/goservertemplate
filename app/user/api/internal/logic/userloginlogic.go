package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"goservertemplate/app/user/api/internal/svc"
	"goservertemplate/app/user/api/internal/types"
	"goservertemplate/common/errorx"
	"time"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *UserLoginLogic) getJwtToken(username string) (string, error) {
	claims := make(jwt.MapClaims)
	var now = time.Now().Unix()
	claims["exp"] = now + l.svcCtx.Config.Auth.AccessExpire
	claims["iat"] = now
	claims["username"] = username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
}
func (l *UserLoginLogic) UserLogin(req *types.UserLogin) (resp *types.UserLoginres, err *errorx.CodeError) {
	token, _ := l.getJwtToken("xu756")
	return &types.UserLoginres{
		Token: token,
	}, nil
}
