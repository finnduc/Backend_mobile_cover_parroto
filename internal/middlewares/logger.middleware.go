package middlewares

import (
	"fmt"
	"go-cover-parroto/global"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)

		if global.Logger != nil {
			fields := []zap.Field{
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("request_id", c.GetString("request_id")),
				zap.String("user_id", fmt.Sprintf("%v", c.GetUint("user_id"))),
				zap.Duration("latency", latency),
			}
			if len(c.Errors) > 0 {
				fields = append(fields, zap.String("errors", c.Errors.String()))
			}

			switch status := c.Writer.Status(); {
			case status >= 500:
				global.Logger.Error("HTTP Request", fields...)
			case status >= 400:
				global.Logger.Warn("HTTP Request", fields...)
			default:
				global.Logger.Info("HTTP Request", fields...)
			}
		}
	}
}