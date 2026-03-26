package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// TokenClaims holds our JWT payload
type TokenClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	JTI    string `json:"jti"` // unique token ID (for blacklist)
	jwt.RegisteredClaims
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "dev-secret-change-me"
	}
	return []byte(secret)
}

// GenerateAccessToken creates JWT access token. Returns (token, jti, error)
func GenerateAccessToken(userID uint, role string, ttl time.Duration) (string, string, error) {
	jti := uuid.NewString()
	claims := TokenClaims{
		UserID: userID,
		Role:   role,
		JTI:    jti,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(getJWTSecret())
	return signed, jti, err
}

// GenerateRefreshToken creates an opaque UUID-based refresh token
func GenerateRefreshToken() string {
	return uuid.NewString()
}

// ParseToken parses and validates JWT, returns claims
func ParseToken(tokenStr string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return getJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}
