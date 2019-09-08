package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	// 创建一个文件
	f, _ := os.Create("12_middleware_gin/gin.log")
	// 配置信息默认输出地址，系统默认为控制台，此时定位为当前路径下日志文件
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	// 创建一个干净的实例
	r := gin.New()

	// 使用中间件
	r.Use(gin.Logger())

	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		c.String(http.StatusOK, "%s", name)
	})
	r.Run()
}
