package configs

import "go-cover-parroto/internal/utils"

type DeepgramConfig struct {
	APIKey string
}

func loadDeepgramConfig() DeepgramConfig {
	return DeepgramConfig{
		APIKey: utils.GetEnv("DEEPGRAM_API_KEY", ""),
	}
}
