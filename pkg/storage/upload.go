package storage

import (
	"github.com/Zmahl/image_recognition_application/pkg/auth"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context, credentials *auth.GoogleCloudCredentials) (string, error) {
	return UploadToGCP(c, credentials)
}
