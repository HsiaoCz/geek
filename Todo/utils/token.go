package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// 这里额外增加一个字段 userId
type myClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// 定义过期时间
const TokenExpireDuration = time.Hour * 24 * 90

// 定义secret

var mySecret = []byte("hello kitty")

// GenToken 生成token
func GenToken(userId int64) (token string, err error) {
	// 创建一个我们自己声明的数据
	claims := myClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "leaf",
		},
	}

	// 使用指定的签名方法创建签名对象
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySecret)
	if err != nil {
		return
	}
	return
}

// ParseToken 解析token
func ParseToken(tokenString string) (*myClaims, error) {
	// 解析自定义的Claim结构体需要使用ParseWithClaims方法
	token, err := jwt.ParseWithClaims(tokenString, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 对token进行类型断言
	if claims, ok := token.Claims.(*myClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
