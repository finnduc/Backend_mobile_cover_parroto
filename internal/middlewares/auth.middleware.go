package middlewares

import (
	"context"
	"fmt"
	"go-familytree/global"
	pkgerrors "go-familytree/pkg/errors"
	"go-familytree/pkg/response"
	"go-familytree/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AuthMiddleware validates JWT and sets user_id + role in context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.ErrorResponseData(c, response.CodeUnauthorized, nil)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Check blacklist
		if global.Rdb != nil {
			claims, err := utils.ParseToken(tokenStr)
			if err == nil {
				blacklistKey := fmt.Sprintf("blacklist:%s", claims.JTI)
				exists, _ := global.Rdb.Exists(context.Background(), blacklistKey).Result()
				if exists > 0 {
					response.ErrorResponseData(c, response.CodeTokenInvalid, nil)
					c.Abort()
					return
				}
			}
		}

		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			global.Logger.Warn("AuthMiddleware: invalid token",
				zap.Error(err),
				zap.String("request_id", c.GetString("request_id")),
			)
			code := response.CodeTokenInvalid
			if err.Error() == pkgerrors.ErrTokenExpired.Error() {
				code = response.CodeTokenExpired
			}
			response.ErrorResponseData(c, code, nil)
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Set("raw_token", tokenStr)
		c.Next()
	}
}