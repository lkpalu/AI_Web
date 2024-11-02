package main

import (
	"Gwe/Gwe"
	"fmt"
	//Gwe "github.com/gin-gonic/gin"
)

type T struct {
	Text string `json:"text"`
}

func main() {
	var TEXT T
	r := Gwe.New()
	r.Use(Gwe.Logger())
	mainGroup := r.Group("/")
	mainGroup.GET("/hello", func(c *Gwe.Context) {
		c.JSON(200, "Hello Gwe")
	})
	mainGroup.GET("/hello/:name", func(c *Gwe.Context) {
		c.JSON(200, "Hello "+c.Param("name"))
		_ = c.ShouldBindJSON(&TEXT)
		fmt.Println(TEXT.Text)
	})
	r.Run(":8080")
}
