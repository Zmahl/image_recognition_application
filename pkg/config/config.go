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
	var config ApplicationConfig

	setCloud(&config, cloudEnv)
	setLabel(&config)

	return config
}

func setCloud(config *ApplicationConfig, cloudEnv string) {
	if cloudEnv == "GCP" {
		config.Storage = storage.CreateGCPStorage(GCP_BUCKET_ENV)
	}

	if cloudEnv == "AWS" {
		config.Storage = storage.AWSProvider{}
	}
}

func setLabel(config *ApplicationConfig) {
	config.Labeller = label.CreateGoogleVision(GOOGLE_VISION_ENV)
}

func GetCloudEnvironment() string {
	return utils.GetEnv("CLOUD_ENV", "")
}
