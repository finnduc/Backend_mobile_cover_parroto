package configs

import "go-cover-parroto/internal/utils"

type FirebaseConfig struct {
	CredentialsFile string
	// ProjectID       string
	// WebAPIKey       string
}

func loadFirebaseConfig() FirebaseConfig {
	return FirebaseConfig{
		CredentialsFile: utils.GetEnv("FIREBASE_CREDENTIALS_FILE", ""),
		// ProjectID:       utils.GetEnv("FIREBASE_PROJECT_ID", ""),
		// WebAPIKey:       utils.GetEnv("FIREBASE_WEB_API_KEY", ""),
	}
}
