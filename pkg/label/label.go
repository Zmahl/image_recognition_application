package label

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/Zmahl/image_recognition_application/pkg/auth"
	"github.com/gin-gonic/gin"
)

func LabelImage(c *gin.Context, credentials *auth.GoogleCloudCredentials, fileName string) {
	var b bytes.Buffer
	uri := fmt.Sprintf("gs://%s/%s", credentials.BucketName, fileName)

	labels, err := detectLabelsURI(&b, uri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	// Returns an array of labels, can be empty
	c.JSON(http.StatusOK, gin.H{
		"labels": labels.LabelAnnotations,
	})
}

func detectLabelsURI(w io.Writer, file string) (*LabelResponse, error) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return &LabelResponse{}, err
	}

	image := vision.NewImageFromURI(file)
	annotations, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		return &LabelResponse{}, err
	}

	var labels LabelResponse

	for _, annotation := range annotations {
		labels.LabelAnnotations = append(labels.LabelAnnotations, annotation.Description)
	}

	return &labels, nil
}
