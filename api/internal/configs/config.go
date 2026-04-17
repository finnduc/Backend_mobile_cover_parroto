package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret        string
	AccessExpiry  int
	RefreshExpiry int
}

func Load() Config {
	_ = godotenv.Load()
	return Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "3001"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Postgres: PostgresConfig{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnvInt("POSTGRES_PORT", 5432),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			DBName:   getEnv("POSTGRES_DB", "parroto"),
			SSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:        getEnv("JWT_SECRET", ""),
			AccessExpiry:  getEnvInt("JWT_ACCESS_EXPIRY", 3600),
			RefreshExpiry: getEnvInt("JWT_REFRESH_EXPIRY", 604800),
		},
	}
}

func getEnv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}

func getEnvInt(key string, def int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return def
}
