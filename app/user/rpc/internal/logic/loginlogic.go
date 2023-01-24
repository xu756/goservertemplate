package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"goservertemplate/app/user/model"
	"log"
	"time"

	"goservertemplate/app/user/rpc/internal/svc"
	"goservertemplate/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *LoginLogic) getJwtToken(username string) (string, error) {
	claims := make(jwt.MapClaims)
	var now = time.Now().Unix()
	claims["exp"] = now + l.svcCtx.Config.Jwt.AccessExpire
	claims["iat"] = now
	claims["username"] = username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(l.svcCtx.Config.Jwt.AccessSecret))
}
func (l *LoginLogic) Login(in *user.User) (*user.LoginResponse, error) {
	token, err := l.getJwtToken(in.Username)
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username:   in.Username,
		Password:   in.Password,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &user.LoginResponse{
		Token: token,
	}, nil
}
