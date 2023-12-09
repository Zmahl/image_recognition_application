package main

import (
	"github.com/Zmahl/image_recognition_application/pkg/config"
	"github.com/Zmahl/image_recognition_application/pkg/label"
	"github.com/Zmahl/image_recognition_application/pkg/storage"
	"github.com/gin-gonic/gin"
)

var conf *config.ApplicationConfig

func init() {
	conf = config.New()
}

func main() {
	r := gin.Default()
	var storage storage.StorageProvider
	var label label.Labeller

	storage = conf.Storage
	label = conf.Labeller

	r.POST("labels/image", label.LabelImageHandler(storage, label))
	r.Run()
}
