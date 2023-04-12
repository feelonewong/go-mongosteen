package controller

import "github.com/gin-gonic/gin"

func PingHandle(c *gin.Context) {
	c.Writer.WriteString("Hello World")
}
