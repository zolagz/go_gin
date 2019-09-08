# 请求路由 - 泛绑定

```$xslt
func main() {
	r := gin.Default()
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	r.Run()
}
```
1. 所有 user 前缀的请求都会映射到这个路由里面，如：
    * http://localhost:8080/user/loedan
    * http://localhost:8080/user/macgrady
2. "/user/*action" 中 *action 表示路由名称的通配符
