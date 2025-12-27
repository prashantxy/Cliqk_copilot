package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/generate/ad-copy", func(c *gin.Context) {
		var req AdRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		prompt := buildPrompt(req)
		response := callLLM(prompt)

		c.JSON(200, response)
	})
	r.Run()
}
