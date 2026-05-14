package configs

import "go-cover-parroto/internal/utils"

type ClerkConfig struct {
	ClerkSecret string
	// ProjectID       string
	// WebAPIKey       string
}

func loadFirebaseConfig() ClerkConfig {
	return ClerkConfig{
		ClerkSecret: utils.GetEnv("CLERK_SECRET", ""),
		// ProjectID:       utils.GetEnv("FIREBASE_PROJECT_ID", ""),
		// WebAPIKey:       utils.GetEnv("FIREBASE_WEB_API_KEY", ""),
	}
}
