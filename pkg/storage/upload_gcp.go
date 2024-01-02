package storage

import (
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"github.com/Zmahl/image_recognition_application/pkg/utils"
	"github.com/gin-gonic/gin"
)

const (
	bucketEnv         = "BUCKET_NAME"
	serviceAccountEnv = "SERVICE_ACCOUNT"
)

type GCPProvider struct {
	BucketName     string
	ServiceAccount string
}

func (gcp GCPProvider) Upload(c *gin.Context) (string, error) {
	ctx := context.Background()

	// This should now be using ADC to access Google Cloud
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer storageClient.Close()

	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		return "", err
	}

	defer f.Close()
	log.Println(gcp.BucketName, uploadedFile.Filename)
	sw := storageClient.Bucket(gcp.BucketName).Object(uploadedFile.Filename).NewWriter(ctx)
	if _, err := io.Copy(sw, f); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		log.Println(err)
		return "", err
	}

	// Attrs() method will return an error if the object does not exist in bucket
	_, err = storageClient.Bucket(gcp.BucketName).Object(uploadedFile.Filename).Attrs(ctx)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("gs://%s/%s", gcp.BucketName, uploadedFile.Filename)

	return url, nil
}

func CreateGCPStorage() *GCPProvider {
	return &GCPProvider{
		BucketName:     utils.GetEnv(bucketEnv, ""),
		ServiceAccount: utils.GetEnv(serviceAccountEnv, ""),
	}
}

func (gcp GCPProvider) GetBucket() string {
	return gcp.BucketName
}

func (gcp GCPProvider) GetServiceAccount() string {
	return gcp.ServiceAccount
}
