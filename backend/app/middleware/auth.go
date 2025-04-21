package middleware

import (

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xiao-en-5970/Goodminton/backend/app/global"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/auth"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/codes"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/responce"
)



func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 从 Header 获取 Token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			responce.SuccessWithCode(c,codes.CodeAuthNotExistError)
			c.Abort()
			return
		}

		// 去掉 "Bearer " 前缀
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

        token, err := auth.ParseToken(tokenString)
        if err != nil {
            global.Logger.Error("Token无效",err)
                
            responce.SuccessWithCode(c,codes.CodeAuthUnvalidToken)
			c.Abort()
            return
        }

        // 安全类型断言（此时已通过ParseToken验证）
        claims := token.Claims.(jwt.MapClaims)
        username := claims["username"].(string)
        
        c.Set("username", username)
        c.Next()
    }
}