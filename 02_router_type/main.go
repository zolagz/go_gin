package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// get 请求
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})
	// post 请求
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})
	// 定义 put delete 等类型请求
	r.DELETE("/delete1", func(c *gin.Context) {
		c.String(200, "delete1")
	})
	// 传参方式定义路由
	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})
	// 路由类型泛解析
	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})
	r.Run()
}