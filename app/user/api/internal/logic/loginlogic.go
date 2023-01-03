package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"goservertemplate/common/errorx"
	"time"

	"goservertemplate/app/user/api/internal/svc"
	"goservertemplate/app/user/api/internal/types"

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

func (l *LoginLogic) getJwtToken(username string) (string, error) {
	claims := make(jwt.MapClaims)
	var now = time.Now().Unix()
	claims["exp"] = now + l.svcCtx.Config.Auth.AccessExpire
	claims["iat"] = now
	claims["username"] = username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
}
func (l *LoginLogic) Login(req *types.Login) (resp *types.Loginres, err *errorx.CodeError) {
	token, _ := l.getJwtToken("xu756")
	return &types.Loginres{
		Token: token,
	}, nil
}
