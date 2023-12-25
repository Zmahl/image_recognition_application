package storage

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type AWSProvider struct {
}

func (aws AWSProvider) Upload(c *gin.Context) (string, error) {
	return "", errors.New("not implemented")
}

func (aws AWSProvider) GetBucket() string {
	return ""
}
