package storage

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/Zmahl/image_recognition_application/pkg/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	"cloud.google.com/go/storage"
)

type GCPProvider struct {
}

func (GCPProvider) Upload(c *gin.Context, credentials *auth.GoogleCloudCredentials) (string, error) {
	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(credentials.CloudStorageServiceAccount))
	if err != nil {
		return "", err
	}

	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		return "", err
	}

	defer f.Close()

	sw := storageClient.Bucket(credentials.BucketName).Object(uploadedFile.Filename).NewWriter(ctx)
	if _, err := io.Copy(sw, f); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
		})
		return "", err
	}

	// checks that the object was uploaded by checking a valid url.
	// sw.Attrs is only valid if there is a successfully written object
	_, err = url.Parse("/" + credentials.BucketName + "/" + sw.Attrs().Name)
	if err != nil {
		return "", err
	}

	return sw.Attrs().Name, nil
}

func (GCPProvider) GetBucketName() {

}

func (GCPProvider) GetCloudCredentials() {

}
