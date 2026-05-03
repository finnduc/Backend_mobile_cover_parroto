package middleware

import (
	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(requiredRole enums.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {

		userRole, ok := c.Request.Context().
			Value(enums.ContextKeyUserRole).(enums.UserRole)

		if !ok {
			logger.S().Errorf("user role not found in context")
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				response.Fail(response.Forbidden("access denied")))
			return
		}

		if userRole != requiredRole {
			logger.S().Errorw("access denied: user role mismatch", "required", requiredRole, "actual", userRole)
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				response.Fail(response.Forbidden("access denied")))
			return
		}

		c.Next()
	}
}
