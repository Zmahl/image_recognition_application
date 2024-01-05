package label

import (
	"net/http"
	"net/http/httptest"
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

func CreateMockContext() *gin.Engine {
	router := gin.Default()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "labels/image", nil)
	router.ServeHTTP(w, req)

	return router
}

func TestVision(t *testing.T) {

}
