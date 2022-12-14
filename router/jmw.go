package router

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var MySecret = []byte("123456") // 定义secret，后面会用到
type MyClaims struct {
	Username             string `json:"username"`
	ID                   int    `json:"id"`
	Role                 int    `json:"role"`
	Ip                   string `json:"ip"`
	jwt.RegisteredClaims        // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

func MakeToken(Username string, Id, Role int, Ip string) (tokenString string, err error) {
	claim := MyClaims{
		Username: Username,
		ID:       Id,
		Role:     Role,
		Ip:       Ip, RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString(MySecret)
	return tokenString, err
}
func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("手写的从前"), nil // 这是我的secret
	}
}

// 解析token
func ParseToken(tokens string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokens, &MyClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("这个token不是一个有效的token1")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("这个token已经过期了")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("这个token还没生效")
			} else {
				return nil, errors.New("这个token不是一个有效的token2")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("这个token不是一个有效的token3")
}

// 中间件jwt
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/api/v1/" {
			token, _ := MakeToken("admin", 1, 1, "127.0.0.1")
			c.Set("jwt", token)
			c.Next()
			return
		}
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(200, gin.H{"code": 400, "msg": "token为空"})
			c.Abort()
			return
		} else {
			claims, err := ParseToken(token) // 解析token
			if err != nil {
				c.JSON(200, gin.H{"code": 400, "msg": err.Error()})
				c.Abort()
				return
			}
			c.Set("jwt", claims)
			c.Next()
		}
	}
}
