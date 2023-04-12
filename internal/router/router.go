package router

import (
	"github.com/gin-gonic/gin"
	"go-mongosteen/internal/controller"
)

func New() *gin.Engine {
	r := gin.Default()
	// ping test
	r.GET("/ping", controller.PingHandle)
	return r
}
