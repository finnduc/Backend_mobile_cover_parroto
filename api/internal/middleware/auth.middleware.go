package middleware

import (
	"context"
	"net/http"
	"strings"

	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/firebase"

	"github.com/gin-gonic/gin"
)

func FirebaseAuth(fbAuth firebase.IFirebaseAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			logger.S().Errorw("invalid authorization header", "header", authHeader)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Fail(response.Unauthorized()))
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		decoded, err := fbAuth.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			logger.S().Errorw("failed to verify id token", "error", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Fail(response.Unauthorized("invalid token")))
			return
		}

		ctx := c.Request.Context()
		userID := decoded.UID
		role, ok := decoded.Claims[string(enums.CustomClaimKeyUserRole)].(enums.UserRole)
		if !ok {
			role = enums.UserRoleGuest
		}

		ctx = context.WithValue(ctx, enums.ContextKeyUserID, userID)
		ctx = context.WithValue(ctx, enums.ContextKeyUserRole, role)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
