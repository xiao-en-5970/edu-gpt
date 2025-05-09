package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
)

var JWTSecret = []byte("your-secret-key")

func GenerateToken(id uint) (string, error) {
	global.Logger.Info(id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(global.Cfg.Auth.MaxAge * time.Hour).Unix(),
	})
	return token.SignedString(JWTSecret)
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	// 1. 解析Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
		}
		return JWTSecret, nil
	})

	// 2. 统一错误处理
	if err != nil {
		return nil, err
	}

	// 3. 验证Token有效性
	if !token.Valid {
		return nil, fmt.Errorf("无效的Token")
	}

	// 4. 验证Claims类型
	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, fmt.Errorf("Claims类型错误")
	}

	return token, nil
}
