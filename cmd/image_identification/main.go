package main

import (
	"github.com/Zmahl/image_recognition_application/pkg/auth"
	"github.com/Zmahl/image_recognition_application/pkg/label"
	"github.com/gin-gonic/gin"
)

var CloudCredentials = auth.CreateCloudCredentials()

func main() {
	r := gin.Default()
	r.POST("labels/image", label.LabelImageHandler(CloudCredentials))
	r.Run()
}
