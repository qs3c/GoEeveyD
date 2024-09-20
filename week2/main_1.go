package main

import (
	"github.com/gin-gonic/gin"
	"week2/testHandlerOfRouter"
)

func main() {
	u := &testHandlerOfRouter.UserHandler{}
	server := gin.Default()
	// 路由注册

	//server.PUT("/users/signup",u.Signup)
	//server.POST("/users/login",u.Login)
	//server.GET("/users/edit",u.Edit)

	// 把路由注册全部写在 main 这边会很多很长
	// 所以打包放到一个函数里 func(* gin.Engine) {...// 注册操作}
	// 这个注册所有路由方法的方法，也绑定到 Handler 结构体上
	// u.RegisterRouters(server)

	// 以 user 进行分组路由
	ug := server.Group("/user")
	u.RegisterRoutersWithGroup(ug)

	_ = server.Run(":8080")
}
