package storage

import (
	"errors"
	"mime/multipart"
)

type AWSProvider struct {
}

func (aws AWSProvider) Upload(file multipart.File, fileName string) (string, error) {
	return "", errors.New("not implemented")
}

func (aws AWSProvider) GetBucket() string {
	return ""
}
