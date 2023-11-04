package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// 定义配置
type myClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// 定义过期时间
const TokenExpireDuration = time.Hour * 24 * 90

// 定义sercet
var mySecret = []byte("xiaofanyi")

// GenToken 生成token
func GenToken(userID int64) (token string, err error) {
	// 创建一个我们自己声明的数据
	claims := myClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "Tenz",
		},
	}
	// 使用指定的签名方式创建签名对象
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySecret)
	if err != nil {
		return
	}
	return token, err
}

// 解析token
func ParseToken(tokenString string) (*myClaims, error) {
	// 解析自定义的claim结构体需要使用ParseWithCliams方法
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
	return nil, errors.New("invaild token")
}
