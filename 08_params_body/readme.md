# 获取请求体的 body 内容

```$xslt
func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		firstName := c.PostForm("first_name")
		lastName  := c.DefaultPostForm("last_name", "last_default_name")
		c.String(http.StatusOK, "%s, %s, %s", firstName, lastName, string(bodyBytes))
	})
	r.Run()
}
```
1. ioutil.ReadAll(c.Request.Body) 将请求内容的 body 读取到一个字节流中
2. c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))，
    由于 ioutil.ReadAll() 方法执行之后，body 数据已经读取到了字节流中，
    此时直接用 c.PostForm("first_name") 将无法拿到对应的数据，
    所以用这个方法将数据回写到 c.Request.Body 中
