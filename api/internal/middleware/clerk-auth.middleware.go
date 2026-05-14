package middleware

import (
	"context"
	"go-cover-parroto/internal/configs"
	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/logger"
	"go-cover-parroto/internal/core/response"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/gin-gonic/gin"
)

type ClerkMetadata struct {
	Role     enums.UserRole `json:"role"`
	Username string         `json:"username"`
}

func customClaimsConstructor(ctx context.Context) any {
	return &ClerkMetadata{}
}

func withCustomClaims(params *clerkhttp.AuthorizationParams) error {
	params.VerifyParams.CustomClaimsConstructor = customClaimsConstructor
	return nil
}

func ClerkAuthMiddleware() gin.HandlerFunc {
	logger := logger.S()
	clerkCfg := configs.Load().ClerkAuth
	clerk.SetKey(clerkCfg.ClerkSecret)

	return func(c *gin.Context) {
		var nextReq *http.Request

		handler := clerkhttp.WithHeaderAuthorization(withCustomClaims)(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextReq = r
			}),
		)

		handler.ServeHTTP(c.Writer, c.Request)

		if nextReq == nil {
			logger.Warnw("Clerk rejected before inner handler",
				"path", c.Request.URL.Path,
				"method", c.Request.Method,
				"authHeader", c.GetHeader("Authorization"),
				"status", c.Writer.Status(),
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Unauthorized())
			return
		}

		claims, ok := clerk.SessionClaimsFromContext(nextReq.Context())
		if !ok || claims.Subject == "" {
			logger.Warnw("Invalid clerk session", "path", nextReq.URL.Path, "method", nextReq.Method, "subject", claims.Subject)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Unauthorized())
			return
		}

		userID := claims.Subject
		role := enums.UserRoleUser

		logger.Warnw("Clerk claims debug",
			"subject", userID,
			"hasCustomClaims", claims.Custom != nil,
		)

		if customClaims, ok := claims.Custom.(*ClerkMetadata); ok && customClaims.Role != "" {
			role = customClaims.Role
		}

		ctx := context.WithValue(nextReq.Context(), enums.ContextKeyUserID, userID)
		ctx = context.WithValue(ctx, enums.ContextKeyUserRole, role)

		c.Request = nextReq.WithContext(ctx)
		c.Next()
	}
}
