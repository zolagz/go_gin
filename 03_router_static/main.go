package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 静态文件夹绑定
	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.Run()
}
