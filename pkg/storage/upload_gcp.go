package storage

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	credentials "cloud.google.com/go/iam/credentials/apiv1"
	"cloud.google.com/go/iam/credentials/apiv1/credentialspb"
	"cloud.google.com/go/storage"
)

type GCPProvider struct {
	BucketName     string
	ServiceAccount string
}

func (gcp GCPProvider) Upload(c *gin.Context) (string, error) {
	ctx := context.Background()
	credCtx, err := credentials.NewIamCredentialsClient(ctx)
	if err != nil {
		return "", err
	}

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

	// Need signed url for labellers to access
	url, err := storage.SignedURL(gcp.BucketName, uploadedFile.Filename, &storage.SignedURLOptions{
		Method:         "GET",
		GoogleAccessID: gcp.ServiceAccount,
		SignBytes: func(b []byte) ([]byte, error) {
			req := &credentialspb.SignBlobRequest{
				Payload: b,
				Name:    gcp.ServiceAccount,
			}
			resp, err := credCtx.SignBlob(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp.SignedBlob, err
		},
		Expires: time.Now().Add(time.Second * 60),
	})
	log.Print(err)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (gcp GCPProvider) GetBucket() string {
	return gcp.BucketName
}

func (gcp GCPProvider) GetServiceAccount() string {
	return gcp.ServiceAccount
}
