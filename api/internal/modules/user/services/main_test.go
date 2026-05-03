package services

import (
	"os"
	"testing"

	"go-cover-parroto/internal/configs"
	"go-cover-parroto/internal/core/logger"
)

func TestMain(m *testing.M) {
	logger.Init(configs.LoggerConfig{Mode: "development"})
	os.Exit(m.Run())
}
