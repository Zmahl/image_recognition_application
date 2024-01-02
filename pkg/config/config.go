package config

import (
	"github.com/Zmahl/image_recognition_application/pkg/label"
	"github.com/Zmahl/image_recognition_application/pkg/storage"
)

type ApplicationConfig struct {
	Storage  storage.StorageProvider
	Labeller label.Labeller
}

func CreateAppConfig(cloudEnv string) ApplicationConfig {
	if cloudEnv == "GCP" {
		return ApplicationConfig{
			Storage:  storage.CreateGCPStorage(),
			Labeller: label.CreateGoogleVision(),
		}
	} else if cloudEnv == "AWS" {
		return ApplicationConfig{
			Storage: storage.AWSProvider{},
		}
	} else {
		return ApplicationConfig{}
	}
}
