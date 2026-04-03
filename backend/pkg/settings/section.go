package settings

type Config struct {
	Server   ServerSetting   `mapstructure:"server"`
	Postgres PostgresSetting `mapstructure:"postgres"`
	Redis    RedisSetting    `mapstructure:"redis"`
	JWT      JWTSetting      `mapstructure:"jwt"`
	Logger   LoggerSetting   `mapstructure:"logger"`
	Scoring  ScoringSetting  `mapstructure:"scoring"`
}

type ServerSetting struct {
	Port         int    `mapstructure:"port"`
	Mode         string `mapstructure:"mode"` // "debug" | "release"
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

type PostgresSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	SSLMode         string `mapstructure:"sslmode"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	TTL      struct {
		TranscriptCache string `mapstructure:"transcript_cache"`
		TokenBlacklist  string `mapstructure:"token_blacklist"`
		RefreshToken    string `mapstructure:"refresh_token"`
	} `mapstructure:"ttl"`
}

type JWTSetting struct {
	AccessTokenTTL  string `mapstructure:"access_token_ttl"`
	RefreshTokenTTL string `mapstructure:"refresh_token_ttl"`
	// Secret loaded from env JWT_SECRET
}

type LoggerSetting struct {
	Level      string `mapstructure:"level"` // "debug" | "info"
	FilePath   string `mapstructure:"file_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type ScoringSetting struct {
	Threshold float64 `mapstructure:"threshold"`
}
