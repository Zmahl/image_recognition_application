package config

import (
	"os"

	"github.com/Zmahl/image_recognition_application/pkg/label"
	"github.com/Zmahl/image_recognition_application/pkg/storage"
)

type ApplicationConfig struct {
	Storage  storage.StorageProvider
	Labeller label.Labeller
}

func New() ApplicationConfig {
	cloud_env := getEnv("CLOUD_ENV", "")
	if cloud_env == "GCP" {
		return ApplicationConfig{
			Storage: storage.GCPProvider{
				BucketName: getEnv("BUCKET_NAME", ""),
			},
			Labeller: label.GoogleVision{
				VisionApiKey: getEnv("VISION_API_KEY", ""),
			},
		}
	} else if cloud_env == "AWS" {
		return ApplicationConfig{
			Storage: storage.AWSProvider{},
		}
	} else {
		return ApplicationConfig{}
	}
}

func setStorage() {

}

// Sets values for keys from env
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
