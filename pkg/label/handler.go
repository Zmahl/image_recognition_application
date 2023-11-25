package label

import (
	"net/http"

	"github.com/Zmahl/image_recognition_application/pkg/auth"
	"github.com/Zmahl/image_recognition_application/pkg/storage"
	"github.com/gin-gonic/gin"
)

// Create a closure for the gin context so that credentials can be passed
func LabelImageHandler(credentials *auth.GoogleCloudCredentials) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		fileName, err := storage.UploadFile(c, credentials)
		if err != nil || len(fileName) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error uploading image",
			})
			return
		}
		LabelImage(c, credentials, fileName)
	}

	return fn
}
