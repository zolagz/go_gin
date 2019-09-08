# 静态目录绑定

```$xslt
func main() {
	r := gin.Default()
	// 静态文件夹绑定
	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.Run()
}
```

本案例需要 build 之后运行二进制文件才能测试效果
```$xslt
cd $GOPATH/src/learn_go_gin/03_router_static
go build -o router_static && ./router_static    // build 并且执行文件
```