package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Zmahl/image_recognition_application/pkg/config"
)

func TestMain(m *testing.M) {
	conf = config.New()
	if conf == (config.ApplicationConfig{}) {
		log.Fatalf("Failed to load config")
	}
	setupRouter()
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestRequestWithoutFile(t *testing.T) {
	router := setupRouter()
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httptest.NewRequest("POST", "/labels/image", nil))
	t.Run("Returns 500 status code", func(t *testing.T) {
		if recorder.Code != 500 {
			t.Error("Expected 500, got ", recorder.Code)
		}
	})
}

func TestRequestWithFile(t *testing.T) {
	var buf bytes.Buffer
	multipartWriter := multipart.NewWriter(&buf)
	file, err := os.Open("github.com/Zmahl/image_recognition_application/test_image/pencil-test.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	formWriter, err := multipartWriter.CreateFormFile("file", "pencil-test.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(formWriter, file); err != nil {
		log.Fatal(err)
	}

	multipartWriter.Close()
	router := setupRouter()
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httptest.NewRequest("POST", "/labels/image", &buf))
	t.Run("Returns 200 status code", func(t *testing.T) {
		if recorder.Code != 200 {
			t.Error("Expected 200, got", recorder.Code)
		}
	})
}
