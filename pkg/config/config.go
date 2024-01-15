package config

import (
	"github.com/Zmahl/image_recognition_application/pkg/label"
	"github.com/Zmahl/image_recognition_application/pkg/storage"
	"github.com/Zmahl/image_recognition_application/pkg/utils"
)

const (
	GCP_BUCKET_ENV    = "BUCKET_NAME"
	GOOGLE_VISION_ENV = "VISION_API_KEY"
)

type ApplicationConfig struct {
	Storage  storage.StorageProvider
	Labeller label.Labeller
}

func (conf ApplicationConfig) VerifyConfig() {
	return
}

func CreateAppConfig(cloudEnv string) ApplicationConfig {
	if cloudEnv == "GCP" {
		return ApplicationConfig{
			Storage:  storage.CreateGCPStorage(GCP_BUCKET_ENV),
			Labeller: label.CreateGoogleVision(GOOGLE_VISION_ENV),
		}
	} else if cloudEnv == "AWS" {
		return ApplicationConfig{
			Storage: storage.AWSProvider{},
		}
	} else {
		return ApplicationConfig{}
	}
}

func GetCloudEnvironment() string {
	return utils.GetEnv("CLOUD_ENV", "")
}
