package storage

import (
	"github.com/gin-gonic/gin"
)

type StorageProvider interface {
	Upload(*gin.Context, string) (string, error)
}

func uploadFile(provider StorageProvider) {
	provider.Upload()
}
