package main

import (
	"net/http"

	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("api/v1")
	api.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server := http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
		// ReadHeaderTimeout: 10,
	}
	// log.Println(server.ListenAndServe())
	server.ListenAndServe()
}
