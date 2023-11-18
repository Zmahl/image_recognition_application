package handlers

import (
	helpers "github.com/Zmahl/image_recognition_application/helpers"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

var storageClient *storage.Client

func IdentifyImageHandler(c *gin.Context) {
	fileName := helpers.UploadFile(c)
	if len(fileName) == 0 {
		return
	}
	helpers.LabelImage(c, fileName)
}
