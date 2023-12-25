package label

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/gin-gonic/gin"
)

type LabelResponse struct {
	LabelAnnotations []string `json:"labelAnnotations"`
}

type GoogleVision struct {
	VisionApiKey string
}

func (gv GoogleVision) LabelImage(c *gin.Context, url string) {
	var b bytes.Buffer
	log.Print(url)
	labels, err := getLabelsFromImage(&b, url)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Returns an array of labels, can be empty
	c.JSON(http.StatusOK, gin.H{
		"labels": labels.LabelAnnotations,
	})
}

func getLabelsFromImage(w io.Writer, file string) (*LabelResponse, error) {
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

func (gv GoogleVision) GetLabelCredentials() string {
	return gv.VisionApiKey
}
