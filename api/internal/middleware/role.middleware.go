package middleware

import (
	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(requiredRole enums.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {

		userRole, ok := c.Request.Context().
			Value(enums.ContextKeyUserRole).(enums.UserRole)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				response.Fail(response.Forbidden("access denied")))
			return
		}

		if userRole != requiredRole {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				response.Fail(response.Forbidden("access denied")))
			return
		}

		c.Next()
	}
}
