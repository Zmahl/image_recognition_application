package config

import (
	"testing"

	"github.com/Zmahl/image_recognition_application/pkg/label"
	"github.com/Zmahl/image_recognition_application/pkg/storage"
)

const gcpEnv = "GCP"

func TestEmptyConfig(t *testing.T) {
	t.Run("empty config due to missing test environment", func(t *testing.T) {
		cloudEnv := ""

		config := CreateAppConfig(cloudEnv)

		if config.Storage != nil {
			t.Errorf("storage should not exist, got %v want %v", config, ApplicationConfig{})
		}

		if config.Labeller != nil {
			t.Errorf("labeller should not exist got %v want %v", config, ApplicationConfig{})
		}

	})
}

func TestMissingStorageFields(t *testing.T) {

	badGCPStorage := storage.CreateGCPStorage("")

	t.Run("config missing storage fields", func(t *testing.T) {

		if badGCPStorage.GetBucket() != "" {
			t.Errorf("storage bucket name should not exist, got %s want %s", badGCPStorage.GetBucket(), "")
		}

	})
}

func TestMissingGoogleVision(t *testing.T) {

	badGoogleVision := label.CreateGoogleVision("")

	t.Run("config missing google vision api key", func(t *testing.T) {

		if badGoogleVision.GetLabelCredentials() != "" {
			t.Errorf("label credentials should not exist, got %s want %s", badGoogleVision.GetLabelCredentials(), "")
		}
	})
}
