package label

import (
	"testing"

	"github.com/gin-gonic/gin"
)

type DummyGoogleVision struct {
}

func (d DummyGoogleVision) LabelImage(c *gin.Context, url string) {

}

func CreateDummyGoogleVision() DummyGoogleVision {
	return DummyGoogleVision{}
}

func TestVision(t *testing.T) {

}
