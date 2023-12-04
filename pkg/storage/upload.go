package storage

import (
	"github.com/gin-gonic/gin"
)

type StorageProvider interface {
	Upload(*gin.Context) (string, error)
	GetBucketName() string
	GetCloudCredentials() string
}

func UploadFile(provider StorageProvider, c *gin.Context) (string, error) {
	return provider.Upload(c)
}
