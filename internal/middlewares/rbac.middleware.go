package middlewares

import (
	"go-familytree/pkg/response"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists || userRole.(string) != role {
			response.ErrorResponseData(c, response.CodeForbidden, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
