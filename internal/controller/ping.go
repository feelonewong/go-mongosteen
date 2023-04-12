package controller

import "github.com/gin-gonic/gin"

func PingHandle(c *gin.Context) {
	// c.Request.xxx 读取请求相关
	// c.JSON() 返回响应相关
	// 抽离controller好处
	// 1.减少router的体积 2.方便进行单元测试
	c.Writer.WriteString("Hello World")
}
