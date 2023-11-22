package helpers

import (
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"

	"cloud.google.com/go/storage"
)

var (
	bucketName             = os.Getenv("BUCKET_NAME")
	cloudBucketCredentials = os.Getenv("CLOUD_BUCKET_SERVICE_ACCOUNT_CREDENTIALS")
)

func UploadFile(c *gin.Context) (string, error) {
	ctx := appengine.NewContext(c.Request)

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(cloudBucketCredentials))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
		})
		return "", err
	}

	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
		})
		return "", err
	}

	defer f.Close()

	sw := storageClient.Bucket(bucketName).Object(uploadedFile.Filename).NewWriter(ctx)
	if _, err := io.Copy(sw, f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
		})
		return "", err
	}

	if err := sw.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
		})
		return "", err
	}

	_, err = url.Parse("/" + bucketName + "/" + sw.Attrs().Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": true,
		})
		return "", err
	}

	return sw.Attrs().Name, nil
}
