package util

import (
	//"github.com/gin-gonic/gin"
	gin "ai_web/Gwe/Gwe"
	"github.com/tmc/langchaingo/llms/ollama"
	"net/http"
)

func CteateModel(c *gin.Context, ModelName string) *ollama.LLM {
	llm, err := ollama.New(ollama.WithModel(ModelName))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return nil
	}
	return llm
}
