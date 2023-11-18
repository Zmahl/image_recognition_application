package handlers

import (
	"net/http"

	helpers "github.com/Zmahl/image_recognition_application/helpers"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

var storageClient *storage.Client

func IdentifyImageHandler(c *gin.Context) {
	fileName, err := helpers.UploadFile(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error uploading image",
		})
		return
	}
	if len(fileName) == 0 {
		return
	}
	helpers.LabelImage(c, fileName)
}
