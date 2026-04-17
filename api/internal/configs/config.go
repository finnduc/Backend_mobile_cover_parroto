package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Firebase FirebaseConfig
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

type FirebaseConfig struct {
	CredentialsFile string
	ProjectID       string
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
		Firebase: FirebaseConfig{
			CredentialsFile: getEnv("FIREBASE_CREDENTIALS_FILE", ""),
			ProjectID:       getEnv("FIREBASE_PROJECT_ID", ""),
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
