# 验证请求参数 - 自定义验证器

```$xslt
// 定义一个接收参数时绑定的结构体
// 在 tag 中定义字段验证规则 binding 属性
// bookabledate 是自定义的验证规则，需要自己实现
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
```

1. 结构体验证的方式采用在结构体的字段 tag 中用 binding 属性来定义验证规则，
    gin 的 ShouldBind 等 bind 方法会自动调用验证规则
2. 验证规则使用 
    * 同时满足用 ',' 分隔 "required,gt=10"
    * 不需要同时满足用 '|' 分隔 "required|gt=10"