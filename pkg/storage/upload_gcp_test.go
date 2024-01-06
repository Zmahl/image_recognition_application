package storage

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"testing"

	"github.com/Zmahl/image_recognition_application/pkg/utils"
)

var testImagePath = utils.GetEnv("TEST_IMAGE", "")

const (
	testBucket = "test_bucket"
	testFile   = "pencil-test"
)

func (d GCPProvider) UploadTestFile(fileName string) (string, error) {
	if fileName == "" {
		return "", errors.New("File name not defined")
	}

	file, fileName, err := createMultipartFile()
	if err != nil {
		return "", err
	}

	url, err := d.Upload(file, fileName)
	if err != nil {
		return "", err
	}

	return url, nil
}

func createMultipartFile() (multipart.File, string, error) {
	var buffer bytes.Buffer

	file, err := os.Open(testImagePath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	w := multipart.NewWriter(&buffer)

	fw, err := w.CreateFormFile("file", testFile)
	if err != nil {
		return nil, "", err
	}

	if _, err := io.Copy(fw, file); err != nil {
		return nil, "", err
	}

}

func CreateTestBucket(bucketName string) GCPProvider {
	return GCPProvider{BucketName: bucketName}
}

func TestUpload(t *testing.T) {
	t.Run("upload should return a google cloud storage string", func(t *testing.T) {
		storage := CreateTestBucket(testBucket)

		got, err := storage.UploadTestFile(testImagePath)
		want := "gs://test_bucket/pencil-test"

		if err != nil {
			t.Errorf("file was empty, got %s want %s", got, want)
		}

		if got != want {
			t.Errorf("url not returned correctly, got %s want %s", got, want)
		}
	})

	t.Run("upload should return a with empty file", func(t *testing.T) {
		storage := CreateTestBucket(testBucket)

		got, err := storage.Upload("")
		want := "gs://test_bucket/"

		if err == nil {
			t.Errorf("file should have been empty, got %s want %s", got, want)
		}
	})

}
