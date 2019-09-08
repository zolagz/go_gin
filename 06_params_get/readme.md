# 获取 get 请求参数

```$xslt
func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName  := c.DefaultQuery("last_name", "last_default_name")
		c.String(http.StatusOK, "%s, %s", firstName, lastName)
	})
	r.Run(":8080")
}
```
1. c.Query("first_name") 获取 get 中的参数
2. c.DefaultQuery("last_name", "last_default_name") 获取 get 的参数，若没有的话则设置为默认值
