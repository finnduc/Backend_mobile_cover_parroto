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

	return Config{
		Server:    loadServerConfig(),
		Postgres:  loadPostgresConfig(),
		ClerkAuth: loadFirebaseConfig(),
		Logger:    loadLoggerConfig(),
	}
}
