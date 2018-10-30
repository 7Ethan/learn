package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // *gin.Engine
	api := router.Group("v1")
	api.GET("ping", func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api.GET("user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Ethan",
		})
	})
	s := http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
