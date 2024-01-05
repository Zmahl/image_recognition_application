package storage

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Zmahl/image_recognition_application/pkg/utils"
)

var testImagePath = utils.GetEnv("TEST_IMAGE", "")

const (
	testBucket = "test_bucket"
	testFile   = "pencil-test"
)

type DummyStorage struct {
	BucketName string
}

func (d DummyStorage) Upload(fileName string) (string, error) {
	if fileName == "" {
		return "", errors.New("File name not defined")
	}

	url := fmt.Sprintf("gs://%s/%s", d.BucketName, testFile)
	return url, nil
}

func (d DummyStorage) GetBucket() string {
	return d.BucketName
}

func CreateDummyBucket(bucketName string) DummyStorage {
	return DummyStorage{BucketName: bucketName}
}

func TestUpload(t *testing.T) {
	t.Run("upload should return a google cloud storage string", func(t *testing.T) {
		storage := CreateDummyBucket(testBucket)

		got, err := storage.Upload(testImagePath)
		want := "gs://test_bucket/pencil-test"

		if err != nil {
			t.Errorf("file was empty, got %s want %s", got, want)
		}

		if got != want {
			t.Errorf("url not returned correctly, got %s want %s", got, want)
		}
	})

	t.Run("upload should return a with empty file", func(t *testing.T) {
		storage := CreateDummyBucket(testBucket)

		got, err := storage.Upload("")
		want := "gs://test_bucket/"

		if err == nil {
			t.Errorf("file should have been empty, got %s want %s", got, want)
		}
	})

}
