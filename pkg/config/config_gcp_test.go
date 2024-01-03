package config

import (
	"testing"

	"github.com/Zmahl/image_recognition_application/pkg/label"
	"github.com/Zmahl/image_recognition_application/pkg/storage"
)

const gcpEnv = "GCP"

type TestApplicationConfig struct {
	Storage  storage.StorageProvider
	Labeller label.Labeller
}

// All tests are initialized without using the GCP environment setup

func TestEmptyConfig(t *testing.T) {
	t.Run("empty config due to missing test environment", func(t *testing.T) {
		cloudEnv := ""

		config := CreateAppConfig(cloudEnv)

		if config.Storage != nil {
			t.Errorf("storage should not exist")
		}

		if config.Labeller != nil {
			t.Errorf("labeller should not exist")
		}

	})
}

func TestIncompleteConfig(t *testing.T) {
	cloudEnv := gcpEnv
	config := CreateAppConfig(cloudEnv)

	t.Run("config missing storage fields", func(t *testing.T) {
		if config.Storage.GetBucket() != "" {
			t.Errorf("storage bucket name should not exist")
		}

		if config.Storage.GetServiceAccount() != "" {
			t.Errorf("storage bucket name should not exist")
		}

	})

	t.Run("config missing labeller fields", func(t *testing.T) {
		if config.Labeller.GetLabelCredentials() != "" {
			t.Errorf("google vision api key should not exist")
		}
	})
}
