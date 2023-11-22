package handlers

import (
	"net/http"

	helpers "github.com/Zmahl/image_recognition_application/helpers"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

var storageClient *storage.Client

func LabelImageHandler(c *gin.Context) {
	fileName, err := helpers.UploadFile(c)
	if err != nil || len(fileName) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error uploading image",
		})
		return
	}

	helpers.LabelImage(c, fileName)
}
