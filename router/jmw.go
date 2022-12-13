package router

import "time"
import "github.com/golang-jwt/jwt/v4"

// 载荷，可添加自己需要的一些信息
type CustomClaims struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
	RoleId   string `json:"roleId"`
	jwt.StandardClaims
}

// TokenExpireDuration token过期时间
const TokenExpireDuration = time.Hour * 24
