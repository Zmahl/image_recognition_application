package storage

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	"github.com/Zmahl/image_recognition_application/pkg/utils"
)

var testImagePath = utils.GetEnv("TEST_IMAGE", "")

const (
	testBucket = "zach-test-bucket-images"
	testFile   = "pencil-test.jpeg"
)

func (d GCPProvider) UploadTestFile(fileName string) (string, error) {
	if fileName == "" {
		return "", errors.New("File name not defined")
	}
	request, err := createRequest()
	if err != nil {
		return "", err
	}

	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		return "", err
	}

	url, err := d.Upload(file, fileHeader.Filename)
	if err != nil {
		return "", err
	}

	return url, nil
}

func CreateTestBucket(bucketName string) GCPProvider {
	return GCPProvider{BucketName: bucketName}
}

func TestUpload(t *testing.T) {

	storage := CreateTestBucket(testBucket)
	t.Run("upload should return a google cloud storage string", func(t *testing.T) {
		got, err := storage.UploadTestFile(testFile)
		want := "gs://zach-test-bucket-images/pencil-test.jpeg"

		if err != nil {
			t.Errorf("file was empty, got %s want %s", got, want)
		}

		if got != want {
			t.Errorf("url not returned correctly, got %s want %s", got, want)
		}
	})

	t.Run("upload should return a with empty file", func(t *testing.T) {
		got, err := storage.UploadTestFile("")
		want := "gs://test_bucket/"

		if err == nil {
			t.Errorf("file should have been empty, got %s want %s", got, want)
		}
	})

	// Teardown code testing the upload
	storage.DeleteImage(testFile)
}

func createRequest() (*http.Request, error) {
	file, err := os.Open(testImagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var buffer bytes.Buffer

	writer := multipart.NewWriter(&buffer)
	part, err := writer.CreateFormFile("file", testFile)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}
	writer.Close()

	bucketURI := fmt.Sprintf("gs://%s/%s", testBucket, testFile)

	req, err := http.NewRequest(http.MethodPost, bucketURI, &buffer)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	client.Do(req)

	return req, err
}
