package middleware

import (
	"ai_web/internal/controller"
	//Nuser "ai_web/internal/user"
	//"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	gin "ai_web/Gwe/Gwe"
)

func Router(r *gin.Engine) gin.HandlerFunc {

	chat := controller.Chat{}
	mainGroup := r.Group("/")
	return func(c *gin.Context) {
		mainGroup.POST("/chat", chat.Dochat)
		//mainGroup.POST("/login", Nuser.Login)
	}
}
