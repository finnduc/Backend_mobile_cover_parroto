package middleware

import (
	"net/http"
	"strings"

	"go-cover-parroto/internal/core/response"
	"go-cover-parroto/internal/database/models"
	fb "go-cover-parroto/internal/firebase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FirebaseAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Fail(response.Unauthorized()))
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		decoded, err := fb.AuthClient.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Fail(response.Unauthorized("invalid token")))
			return
		}

		email, _ := decoded.Claims["email"].(string)

		var user models.User
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Fail(response.Unauthorized("user not found")))
			return
		}

		c.Set("userID", user.ID)
		c.Next()
	}
}
