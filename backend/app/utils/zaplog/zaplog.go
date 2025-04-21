package zaplog

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)



func ZapLogger(sugar *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 处理请求
		c.Next()

		// 记录日志
		sugar.Infow("HTTP Request",
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"ip", c.ClientIP(),
			"latency", time.Since(start),
			"user-agent", c.Request.UserAgent(),
		)
	}
}