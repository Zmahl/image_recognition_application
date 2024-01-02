package config

import "testing"

const gcpEnv = "GCP"

func CreateTestConfig(cloudEnv string) ApplicationConfig {
	return ApplicationConfig{}
}

func TestEmptyConfig(t *testing.T) {
	t.Run("empty config has empty fields and returns error", func(t *testing.T) {
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
	t.Run("empty config has empty fields and returns error", func(t *testing.T) {
		cloudEnv := gcpEnv

		config := CreateAppConfig(cloudEnv)

		if config.Storage != nil {
			t.Errorf("storage should not exist")
		}

		if config.Labeller != nil {
			t.Errorf("labeller should not exist")
		}

	})
}
