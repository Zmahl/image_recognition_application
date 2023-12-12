package label

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func TestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}

func MockLabelImage(c *gin.Context) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "multipart/form-data")
}
