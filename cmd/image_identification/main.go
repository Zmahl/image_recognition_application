package main

import (
	"log"

	"github.com/Zmahl/image_recognition_application/pkg/config"
	"github.com/Zmahl/image_recognition_application/pkg/label"
	"github.com/Zmahl/image_recognition_application/pkg/storage"
	"github.com/gin-gonic/gin"
)

var conf config.ApplicationConfig

func init() {
	conf = config.CreateAppConfig()
	if conf == (config.ApplicationConfig{}) {
		log.Fatalf("Failed to load config")
	}
}

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	var stor storage.StorageProvider
	var lab label.Labeller

	stor = conf.Storage
	lab = conf.Labeller

	router.POST("labels/image", label.LabelImageHandler(stor, lab))
	return router
}
