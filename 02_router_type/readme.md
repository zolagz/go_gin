# 讲解 gin 实现 HTTP 的多种路由实现方式

### 初始化 gin 实例
```$xslt
r := gin.Default()
```

### GET 请求
```$xslt
// get 请求
r.GET("/get", func(c *gin.Context) {
	c.String(200, "get")
})
```

### POST 请求
```$xslt
// post 请求
r.POST("/post", func(c *gin.Context) {
	c.String(200, "post")
})
```

### DELETE 请求
```$xslt
// 定义 put delete 等类型请求
r.DELETE("/delete", func(c *gin.Context) {
	c.String(200, "delete")
})
```

其余如 put，patch 等 HTTP 请求方式的实现类同上面，就不多赘述...

## 路由泛解析的两种方式

* 利用 gin 的 Handle() 方法定义路由
```$xslt
// 传参方式定义路由
r.Handle("DELETE", "/delete_1", func(c *gin.Context) {
	c.String(200, "delete_1")
})
```

* 使用 ANY 的方式定义路由
```$xslt
// 路由类型泛解析
r.Any("/any", func(c *gin.Context) {
	c.String(200, "any")
})
```
