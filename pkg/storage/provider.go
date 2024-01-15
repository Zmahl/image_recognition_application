package storage

import "mime/multipart"

type StorageProvider interface {
	Upload(multipart.File, string) (string, error)
	GetBucket() string
}
