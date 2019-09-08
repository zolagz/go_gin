# gin 中间件

```$xslt
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
```

1. os.Create("12_middleware_gin/gin.log") 创建一个文件
2. 一般创建 gin 实例会用 gin.Default() 方法，方法中默认实现了 Logger() 和 Recover() 两个中间件，
    这里直接 New() 一个干净的实例，然后手动调用中间件调试