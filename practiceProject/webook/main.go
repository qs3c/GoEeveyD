package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"practiceProject/webook/internel/web"
	"strings"
	"time"
)

func main() {

	server := gin.Default()

	//
	u := &web.UserHandler{}

	// 预编译正则表达式子
	u.PreCompile()
	// 注册路由
	u.RegisterRouters(server)

	// Use 中间件与跨域请求解决
	server.Use(func(context *gin.Context) {
		println("中间件1")
	})

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// 是否允许 cookie session 之类的
		AllowCredentials: true,
		// 规则比较复杂的时候
		// 对来源网址 origin 做分析来判断是否可以请求现在的地址
		AllowOriginFunc: func(origin string) bool {
			// 开发环境和来自公司的域名可以访问
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "company.com")
		},
		// preflight 请求的缓存时间
		MaxAge: 12 * time.Hour,
	}))

	_ = server.Run(":8080")
}
