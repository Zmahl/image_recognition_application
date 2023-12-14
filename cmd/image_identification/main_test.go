package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLabelRoute(t *testing.T)  {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/labels/image", func (c *gin.Context)) {
		c.JSON(200, gin.H{

		})
	}
}

w := httptest.NewRecorder()
req, _ := http.NewRequest("POST", "/labels/image", nil)
router.ServeHTTP(w, req)