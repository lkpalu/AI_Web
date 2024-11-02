package middleware

import (
	"github.com/gin-gonic/gin"
	//gin "ai_web/Gwe/Gwe"
	"log"
	"net/http"
)

var origins = []string{
	"https://localhost:8081",
	"https://localhost:8082",
	"https://localhost:63342",
}

//	func Cors() gin.HandlerFunc {
//		return func(c *gin.Context) {
//			origin := c.GetHeader("Origin")
//			log.Println(origin)
//			for _, o := range origins {
//				if o == origin {
//					c.Header("Access-Control-Allow-Origin", o)
//					c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
//					c.Header("Access-Control-Allow-Headers", "content-type")
//
//					if c.Request.Method == "OPTIONS" {
//						c.JSON(http.StatusOK, "")
//						c.Abort()
//						return
//					}
//					c.Next()
//				}
//			}
//		}
//	}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		log.Println(origin)
		//for _, o := range origins {
		//	if o == origin {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "content-type")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "")
			c.Abort()
			return
		}
		c.Next()
	}
	//	}
	//}
}
