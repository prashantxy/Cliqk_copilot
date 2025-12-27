package main

import (
	"cliqk-copilot/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Marketing Copilot API running 🚀",
		})
	})

	r.POST("/generate/ad-copy", handlers.GenerateAdCopy)

	r.Run(":8080")
}
