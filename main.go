package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Image struct {
	ImageData string `json:"data"`
}

type Body struct {
	Name string `json:"name"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})

	r.POST("/identify_image", func(c *gin.Context) {
		body := Image{}
		if err:=c.BindJSON(&body);err!=nil{
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
		c.JSON(http.StatusAccepted, &body)
	})

	r.POST("/test", func(c *gin.Context) {
		body:=Body{}
		if err:=c.BindJSON(&body);err!=nil{
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
		c.JSON(http.StatusAccepted, &body)
	})
	r.Run(":8000")
}