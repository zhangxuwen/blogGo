package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var mySercet = []byte("zhangmuyu")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, usernname string) (string, error) {
	// 创建一个自己的声明
	c := MyClaims{
		UserID:   userID, // 自定义字段
		Username: usernname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt("auth.jwt_expire") * int(time.Hour))).Unix(),
			Issuer: "zhangmuyu",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// if token == nil {
	// 	zap.L().Error("token is null")
	// }
	strs, err := token.SignedString(mySercet)
	if err != nil {
		zap.L().Error("token.SignedString failed", zap.Error(err))
		zap.L().Info(strs)
	}
	return strs, err
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySercet, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		// 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
