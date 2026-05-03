package configs

import "github.com/joho/godotenv"

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Firebase FirebaseConfig
	Logger   LoggerConfig
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		Server:   loadServerConfig(),
		Postgres: loadPostgresConfig(),
		Firebase: loadFirebaseConfig(),
		Logger:   loadLoggerConfig(),
	}
}
