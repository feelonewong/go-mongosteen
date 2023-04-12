package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// ping test
	r.GET("/ping", PingFunc)
	return r
}

func PingFunc(c *gin.Context) {
	c.Writer.WriteString("Hello World")
}
