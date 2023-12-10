package label

import (
	"net/http"

	"github.com/Zmahl/image_recognition_application/pkg/storage"
	"github.com/gin-gonic/gin"
)

// Create a closure for the gin context so that credentials can be passed
func LabelImageHandler(sp storage.StorageProvider, lb Labeller) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		url, err := sp.Upload(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error uploading image",
			})
			return
		}
		lb.LabelImage(c, url)
	}

	return fn
}
