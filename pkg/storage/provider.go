package storage

import (
	"github.com/gin-gonic/gin"
)

type StorageProvider interface {
	Upload(*gin.Context) (string, error)
	GetBucket() string
}
