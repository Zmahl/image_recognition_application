package storage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"github.com/Zmahl/image_recognition_application/pkg/utils"
)

const (
	bucketEnv = "BUCKET_NAME"
)

type GCPProvider struct {
	BucketName string
}

func (gcp GCPProvider) Upload(file multipart.File, fileName string) (string, error) {
	ctx := context.Background()

	// This should now be using ADC to access Google Cloud
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer storageClient.Close()

	sw := storageClient.Bucket(gcp.GetBucket()).Object(fileName).NewWriter(ctx)
	if _, err := io.Copy(sw, file); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}

	// Attrs() method will return an error if the object does not exist in bucket
	_, err = storageClient.Bucket(gcp.GetBucket()).Object(fileName).Attrs(ctx)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("gs://%s/%s", gcp.GetBucket(), fileName)

	return url, nil
}

func (gcp GCPProvider) DeleteImage(fileName string) error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := client.Bucket(gcp.GetBucket()).Object(fileName)

	attrs, err := o.Attrs(ctx)
	if err != nil {
		return err
	}
	o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	if err := o.Delete(ctx); err != nil {
		return err
	}

	return nil
}

func CreateGCPStorage(bucketEnv string) *GCPProvider {
	return &GCPProvider{
		BucketName: utils.GetEnv(bucketEnv, ""),
	}
}

func (gcp GCPProvider) GetBucket() string {
	return gcp.BucketName
}
