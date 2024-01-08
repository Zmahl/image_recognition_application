package label

import (
	"context"
	"io"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/Zmahl/image_recognition_application/pkg/utils"
)

type LabelResponse struct {
	LabelAnnotations []string `json:"labelAnnotations"`
}

type GoogleVision struct {
	VisionApiKey string
}

func (gv GoogleVision) LabelImage(w io.Writer, file string) (*LabelResponse, error) {
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

func CreateGoogleVision(visionApiEnv string) *GoogleVision {
	return &GoogleVision{
		VisionApiKey: utils.GetEnv(visionApiEnv, ""),
	}
}

func (gv GoogleVision) GetLabelCredentials() string {
	return gv.VisionApiKey
}
