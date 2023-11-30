package auth

import (
	"github.com/Zmahl/image_recognition_application/pkg/config"
)

func setGCPCredentials(conf *config.Config) *GoogleCloudCredentials {
	return &GoogleCloudCredentials{
		BucketName:                 conf.Storage.BucketName,
		CloudStorageServiceAccount: conf.Storage.GoogleServiceAccount,
		VisionApiKey:               conf.Vision.VisionAPIKey,
	}
}

func setCredentials(conf *config.Config) *GoogleCloudCredentials {
	credentials := setGCPCredentials(conf)
	return credentials
}

func NewCloudCredentials(conf *config.Config) *GoogleCloudCredentials {
	CloudCredentials := setCredentials(conf)
	return CloudCredentials
}
