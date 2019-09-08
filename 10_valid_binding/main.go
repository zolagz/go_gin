package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义一个接收参数时绑定的结构体
// 在 tag 中定义字段验证规则 binding 属性
type Person struct {
	Age 		string		`form:"age" binding:"required,gt=10"`
	Name 		string		`form:"name" binding:"required"`
	Address 	string		`form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		var person Person
		// 这里是根据请求的 content-type 来做不同的 binding 操作
		// 绑定的结果是将参数绑定到 Person 结构体中
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusBadRequest, "person bind err: %v", err)
			c.Abort()
			return
		} else {
			c.String(http.StatusOK, "%v", person)
		}
	})
	r.Run()
}
