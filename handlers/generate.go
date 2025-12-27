package handlers

import (
	"net/http"

	"cliqk-copilot/models"
	"cliqk-copilot/services"
	"github.com/gin-gonic/gin"
)

func GenerateAdCopy(c *gin.Context) {
	var req models.AdRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	prompt := services.BuildAdPrompt(req)
	result := services.CallLLM(prompt)

	c.JSON(http.StatusOK, result)
}
