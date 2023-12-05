package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Zmahl/image_recognition_application/pkg/auth"
	"github.com/Zmahl/image_recognition_application/pkg/config"
	"github.com/Zmahl/image_recognition_application/pkg/label"
	"github.com/gin-gonic/gin"
)

var conf *config.Config

//Reverted to basic manual check

func validateConfig(conf *config.Config) error {
	// Need this to check for null string pointer values
	invalid_value := ""

	if &conf.Storage.BucketName == &invalid_value {
		return fmt.Errorf("BUCKET_NAME not defined")
	}
	if &conf.Storage.GoogleServiceAccount == &invalid_value {
		return fmt.Errorf("GOOGLE_SERVICE_ACCOUNT not defined")
	}
	if &conf.Vision.VisionAPIKey == &invalid_value {
		return fmt.Errorf("VISION_API_KEY not defined")
	}

	return nil
}

func init() {

	// If there is a value not defined in the Config, app will not start
	conf = config.New()
	err := validateConfig(conf)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}

func main() {
	auth := auth.NewCloudCredentials(conf)
	r := gin.Default()
	storage := GetStorageProvider()
	r.POST("labels/image", label.LabelImageHandler(auth))
	r.Run()
}
