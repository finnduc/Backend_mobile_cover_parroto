package configs

import "go-cover-parroto/internal/utils"

type PostgresConfig struct {
	Host           string
	Port           int
	User           string
	Password       string
	DBName         string
	SSLMode        string
	ChannelBinding string
}

func loadPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:           utils.GetEnv("POSTGRES_HOST", "localhost"),
		Port:           utils.GetEnvInt("POSTGRES_PORT", 5432),
		User:           utils.GetEnv("POSTGRES_USER", "postgres"),
		Password:       utils.GetEnv("POSTGRES_PASSWORD", ""),
		DBName:         utils.GetEnv("POSTGRES_DB", "engflix"),
		SSLMode:        utils.GetEnv("POSTGRES_SSLMODE", "disable"),
		ChannelBinding: utils.GetEnv("POSTGRES_CHANNELBINDING", "*"),
	}
}
