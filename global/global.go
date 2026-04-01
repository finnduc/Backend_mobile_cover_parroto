package global

import (
	"go-cover-parroto/pkg/settings"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *zap.Logger
	DB     *gorm.DB
	Rdb    *redis.Client
)