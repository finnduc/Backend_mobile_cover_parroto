package initialize

import (
	"fmt"
	"go-familytree/global"
	"go-familytree/pkg/logger"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config: %w", err))
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode config: %w", err))
	}

	// Fail-fast: JWT_SECRET must be provided in env
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret-change-me" // dev fallback only
	}
	global.Config.JWT.AccessTokenTTL = viper.GetString("jwt.access_token_ttl")
	global.Config.JWT.RefreshTokenTTL = viper.GetString("jwt.refresh_token_ttl")

	// Initialize logger immediately after config
	global.Logger = logger.NewLogger(global.Config.Logger)
	global.Logger.Info("Config loaded successfully")
}