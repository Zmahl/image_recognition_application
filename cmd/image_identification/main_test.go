package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"

	"github.com/Zmahl/image_recognition_application/pkg/config"
)

func init() {
	conf = config.New()
	if conf == (config.ApplicationConfig{}) {
		log.Fatalf("Failed to load config")
	}
}

func TestLabelRoute(t *testing.T) {
	imageBytes, err := ioutil.ReadFile("github.com/Zmahl/image_recognition_application/pkg/test_image/pencil_test.jpeg")

	router := setupRouter()
	w := httptest.NewRecorder()
	file, err := 
	req, _ := http.NewRequest("POST", "/labels/image", nil)
	router.ServeHTTP(w, req)


}
