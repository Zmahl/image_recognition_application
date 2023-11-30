package main

import (
	"log"
	"os"

	"github.com/Zmahl/image_recognition_application/pkg/config"
	"github.com/Zmahl/image_recognition_application/pkg/label"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var conf *config.Config

func init() {

	// If there is no .env file to load, the program will exit
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		os.Exit(1)
	}
}

func main() {

	r := gin.Default()
	r.POST("labels/image", label.LabelImageHandler(conf))
	r.Run()
}
