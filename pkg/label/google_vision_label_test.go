package label

import (
	"os"
	"testing"
)

func CreateTestGoogleVision() *GoogleVision {
	return &GoogleVision{VisionApiKey: os.Getenv("VISION_API_KEY")}
}

func TestVision(t *testing.T) {
	//googleVision := CreateTestGoogleVision()

	t.Run("testing that the labeller returns the correct annotations from test image", func(t *testing.T) {

	})
}
