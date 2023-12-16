package main

import (
	"log"
	"net/http/httptest"
	"testing"

	"github.com/Zmahl/image_recognition_application/pkg/config"
)

func init() {
	conf = config.New()
	if conf == (config.ApplicationConfig{}) {
		log.Fatalf("Failed to load config")
	}
}

func TestMain(t *testing.T) {
	router := setupRouter()
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httptest.NewRequest("POST", "/labels/image", nil))
	t.Run("Returns 500 status code", func(t *testing.T) {
		if recorder.Code != 500 {
			t.Error("Expected 500, got ", recorder.Code)
		}
	})
}
