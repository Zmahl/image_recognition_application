package label

import (
	"testing"

	"github.com/Zmahl/image_recognition_application/pkg/label"
	"github.com/Zmahl/image_recognition_application/pkg/storage"
	"github.com/gin-gonic/gin"
)

func TestGinContext(m *testing.M) *gin.Context {
	gin.SetMode(gin.TestMode)

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	var stor storage.StorageProvider
	var lab label.Labeller

	stor = conf.Storage
	lab = conf.Labeller

	r.POST("labels/image", label.LabelImageHandler(stor, lab))

	return router
}

func MockLabelImage(c *gin.Context) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "multipart/form-data")
}
