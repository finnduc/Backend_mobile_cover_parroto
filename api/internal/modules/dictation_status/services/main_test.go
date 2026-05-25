package services

import (
	"context"
	"os"
	"testing"

	"go-cover-parroto/internal/configs"
	"go-cover-parroto/internal/core/enums"
	"go-cover-parroto/internal/core/logger"
)

func TestMain(m *testing.M) {
	logger.Init(configs.LoggerConfig{Mode: "development"})
	os.Exit(m.Run())
}

func testCtx(userID string, role enums.UserRole) context.Context {
	ctx := context.WithValue(context.Background(), enums.ContextKeyUserID, userID)
	ctx = context.WithValue(ctx, enums.ContextKeyUserRole, role)
	return ctx
}
