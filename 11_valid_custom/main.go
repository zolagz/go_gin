package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

// 定义一个接收参数时绑定的结构体
// 在 tag 中定义字段验证规则 binding 属性
type Booking struct {
	CheckIn 	time.Time	`form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut 	time.Time	`form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

// 自定义验证器
func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	date, ok := field.Interface().(time.Time)
	if !ok {
		return false
	}
	today := time.Now()
	if date.Unix() > today.Unix() {
		return true
	}
	return false
}

func main() {
	r := gin.Default()

	// 注册自定义的验证器以供使用
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		// 这里是根据请求的 content-type 来做不同的 binding 操作
		// 绑定的结果是将参数绑定到 Person 结构体中
		if err := c.ShouldBind(&b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok!",
		})
	})
	r.Run()
}
