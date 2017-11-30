package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"log"
)

func main() {
	r := gin.Default()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		log.Printf("error %s", "testing")

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
