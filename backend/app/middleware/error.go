package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/edu-gpt/backend/app/utils/responce"
	"go.uber.org/zap"
)

// ErrorMiddleware 自定义错误中间件
func ErrorMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 先执行后续的 Handler
        c.Next() // 执行后续的 Handler，如果有错误会存到 c.Errors

        // 检查是否有错误
        if len(c.Errors) > 0 {
            err := c.Errors.Last() // 获取最后一个错误

            // 用 Zap 记录错误日志（包含请求信息）
            logger.Error("Handler error",
                zap.Error(err.Err),          // 错误本身
                zap.String("path", c.Request.URL.Path),      // 请求路径
                zap.String("method", c.Request.Method),      // 请求方法
                zap.String("client_ip", c.ClientIP()),       // 客户端 IP
            )

            // 统一返回错误响应
            responce.ErrorInternalServerError(c, err.Err)
        }
    }
}
