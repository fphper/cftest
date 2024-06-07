package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建一个默认的Gin引擎
	server := gin.Default()

	// 2. 定义路由：当使用 GET 方法访问路径为 /hello 时，执行回调函数
	server.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Cloudflare Page!")
	})

	// 4. 启动服务，监听在 0.0.0.0:8080 上
	server.Run(":80") // 如果不指定端口号，默认为8080
}
