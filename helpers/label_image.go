package helpers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	types "github.com/Zmahl/image_recognition_application/types"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/gin-gonic/gin"
)

var VISION_API_KEY = os.Getenv("vision_api_key")

func LabelImage(c *gin.Context, fileName string) {
	var b bytes.Buffer
	uri := fmt.Sprint("gs://%w/%w", BUCKET_NAME, fileName)

	labels, err := detectLabelsURI(&b, uri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	if len(labels.LabelAnnotations) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No labels found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"labels": labels.LabelAnnotations,
	})
}

func detectLabelsURI(w io.Writer, file string) (*types.LabelResponse, error) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return &types.LabelResponse{}, err
	}

	image := vision.NewImageFromURI(file)
	annotations, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		return &types.LabelResponse{}, err
	}

	var labels types.LabelResponse

	for _, annotation := range annotations {
		labels.LabelAnnotations = append(labels.LabelAnnotations, annotation.Description)
	}

	return &labels, nil
}
