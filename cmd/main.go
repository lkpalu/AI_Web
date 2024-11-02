package main

import (
	"ai_web/internal/middleware"
	//"github.com/gin-gonic/gin"
	gin "ai_web/Gwe/Gwe"
)

func main() {
	r := gin.New()

	//r.Use(middleware.Cors())
	r.Use(middleware.Router(r))
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
