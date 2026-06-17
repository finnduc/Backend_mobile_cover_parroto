package configs

import "github.com/joho/godotenv"

type Config struct {
	Server    ServerConfig
	Postgres  PostgresConfig
	ClerkAuth ClerkConfig
	Logger    LoggerConfig
}

func Load() Config {
	_ = godotenv.Load()

	cfg := Config{
		Server:    loadServerConfig(),
		Postgres:  loadPostgresConfig(),
		ClerkAuth: loadClerkConfig(),
		Logger:    loadLoggerConfig(),
	}
	return cfg
}
