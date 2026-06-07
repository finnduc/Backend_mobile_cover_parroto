package configs

import "go-cover-parroto/internal/utils"

type AzureConfig struct {
	SpeechKey    string
	SpeechRegion string
}

func LoadAzureConfig() AzureConfig {
	return AzureConfig{
		SpeechKey:    utils.GetEnv("AZURE_SPEECH_KEY", ""),
		SpeechRegion: utils.GetEnv("AZURE_SPEECH_REGION", ""),
	}
}
