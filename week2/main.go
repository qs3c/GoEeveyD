package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 返回一个服务器
	server := gin.Default()
	// 简单定义一个路由注册 <path,func>
	// 200 也可以用 http 里面的 statusOK
	server.GET("/hello", func(c *gin.Context) {
		//c.String(200, "Hello World")
		c.String(http.StatusOK, "Hello World")
	})
	// 启动服务器
	err := server.Run(":8080")
	// 错误处理
	if err != nil {
		return
	}

	// 是一种逻辑上的服务器，可以启动多个
	// 不过这段代码放在这里，可能会应为上面的 Run 阻塞到而执行不到
	// 所以正常来说要放到上面的 Run 之上
	go func() {
		server1 := gin.Default()
		_ = server1.Run(":8081")
	}()

}
