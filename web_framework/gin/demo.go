package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	Warnning *log.Logger
)

func init() {
	file, err := os.OpenFile("demo.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	Warnning = log.New(io.MultiWriter(file, os.Stderr), "Gin demo :",
		log.Ltime|log.Lshortfile)
}

func main() {
	r := gin.Default()
	r.HEAD("keep-alive")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
		c.Header("keep-alive", "0")
	})
	Warnning.Println(r.Run(":8080"))
}
