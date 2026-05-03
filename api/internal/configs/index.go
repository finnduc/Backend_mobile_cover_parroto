package configs

import "github.com/joho/godotenv"

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Firebase FirebaseConfig
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		Server:   loadServerConfig(),
		Postgres: loadPostgresConfig(),
		Firebase: loadFirebaseConfig(),
	}
}
