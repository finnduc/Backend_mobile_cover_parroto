package configs

import "go-cover-parroto/internal/utils"

type LoggerConfig struct {
	Mode string
}

func loadLoggerConfig() LoggerConfig {
	return LoggerConfig{
		Mode: utils.GetEnv("LOGGER_MODE", "development"),
	}
}
