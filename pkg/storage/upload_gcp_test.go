package storage

import (
	"fmt"
	"testing"

	"github.com/Zmahl/image_recognition_application/pkg/utils"
)

var testImagePath = utils.GetEnv("TEST_IMAGE", "")

type DummyStorage struct {
	BucketName string
}

func (d DummyStorage) Upload(fileName string) (string, error) {
	url := fmt.Sprintf("gs://%s/%s", d.BucketName, fileName)
	return url, nil
}

func (d DummyStorage) GetBucket() string {
	return d.BucketName
}

func TestUpload(t *testing.T) {
	fmt.Println(testImagePath)
}
