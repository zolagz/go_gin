# 获取 post 请求参数

```$xslt
func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		firstName := c.PostForm("first_name")
		lastName  := c.DefaultPostForm("last_name", "last_default_name")
		c.String(http.StatusOK, "%s, %s", firstName, lastName)
	})
	r.Run()
}
```
1. c.PostForm("first_name") 获取 post 中的参数
2. c.DefaultPostForm("last_name", "last_default_name") 获取 post 的参数，若没有的话则设置为默认值
