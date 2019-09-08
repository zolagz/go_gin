package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	// 两个路由同时指向一个回调方法
	r.GET("/testing", bindFunc)
	r.POST("/testing", bindFunc)
	r.Run()
}

// 定义一个接收参数时绑定的结构体
type Person struct {
	Name 		string		`form:"name"`
	Address 	string		`form:"address"`
	Birthday 	time.Time	`form:"birthday" time_format:"2006-01-02"`
}

// 路由回调
func bindFunc(c *gin.Context) {
	var person Person
	// 这里是根据请求的 content-type 来做不同的 binding 操作
	// 绑定的结果是将参数绑定到 Person 结构体中
	if err := c.ShouldBind(&person); err != nil {
		c.String(http.StatusOK, "person bind err: %v", err)
	} else {
		c.String(http.StatusOK, "%v", person)
	}
}
