package main

import (
	handlers "github.com/Zmahl/image_recognition_application/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("labels/image", handlers.LabelImageHandler)
	r.Run()
}
