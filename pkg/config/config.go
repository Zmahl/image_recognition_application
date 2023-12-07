package config

import (
	"os"
)

type GoogleCloudStorageConfig struct {
	BucketName           string
	GoogleServiceAccount string
}

type VisionConfig struct {
	VisionAPIKey string
}

type AWSCloudStorageConfig struct {
}

type Config struct {
	Storage GoogleCloudStorageConfig
	Vision  VisionConfig
}

func New() *Config {
	return &Config{
		Storage: GoogleCloudStorageConfig{
			BucketName:           getEnv("BUCKET_NAME", ""),
			GoogleServiceAccount: getEnv("GOOGLE_CLOUD_SERVICE_ACCOUNT", ""),
		},
		Vision: VisionConfig{
			VisionAPIKey: getEnv("VISION_API_KEY", ""),
		},
	}
}

// Sets values for keys from env
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
