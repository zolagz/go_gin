# 演示 gin 的启动案例

```$xslt
package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个 gin 的实例
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()	// 默认端口 8080
}
```

若要使用的端口已被占用，如默认的 8080 端口
```$xslt
lsof -i:8080    // 获取进程 ID
kill ID         // 杀死进程
```

``r.Run()	// 启动服务，默认端口 8080``