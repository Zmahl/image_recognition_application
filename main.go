package main

import (
	handlers "github.com/Zmahl/image_recognition_application/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("label-image", handlers.IdentifyImageHandler)
	r.Run()
}
