package label

import (
	"net/http"

	"github.com/Zmahl/image_recognition_application/pkg/storage"
	"github.com/gin-gonic/gin"
)

// Create a closure for the gin context so that credentials can be passed
func LabelImageHandler(sp storage.StorageProvider, lb Labeller) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		file, fileHeader, err := c.Request.FormFile("file")
		url, err := sp.Upload(file, fileHeader.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error uploading image",
			})
			return
		}

		labels, err := lb.LabelImage(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error identifying image",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"labels": &labels,
		})
	}

	return fn
}
