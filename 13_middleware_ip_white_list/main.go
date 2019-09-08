package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 自定义中间件
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIP := c.ClientIP()
		for _, ip := range ipList {
			if ip == clientIP {
				flag = true
			}
		}
		if !flag {
			c.String(http.StatusUnauthorized, "%s not in ipList", clientIP)
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()

	// 使用自定义的中间件
	r.Use(IPAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_name")
		c.String(http.StatusOK, "%s", name)
	})
	r.Run()
}
