# 参数作为 uri 功能

```$xslt
func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": c.Param("name"),
			"id"  : c.Param("id"),
		})
	})
	r.Run()
}
```
1. 路由中 /:name/:id 使用 RestFul 的方式将参数绑定到 uri 中
2. c.Param() 可以获取 uri 中的参数
