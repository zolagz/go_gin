package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个 gin 的实例
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()	// 启动服务，默认端口 8080
}
