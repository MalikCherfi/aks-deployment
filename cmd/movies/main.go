package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/movies", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app":  "Movies API",
			"data": []string{"Movie 1", "Movie 2", "Movie 3"},
		})
	})

	r.GET("/healthz", func(c *gin.Context) {
    	c.JSON(http.StatusOK, gin.H{
        "status": "ok",
    	})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
