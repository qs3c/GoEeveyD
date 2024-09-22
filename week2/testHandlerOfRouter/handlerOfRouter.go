package testHandlerOfRouter

import "github.com/gin-gonic/gin"

// 把所有和某个功能下的所有路由方法绑定到一个结构体上（包括注册这些方法的方法也绑定到这个结构体上）
// 比如：用户注册、登陆等相关的 “路由方法” 都绑定到 “用户” 结构体上

type UserHandler struct{}

// 路由方法：

func (u *UserHandler) Signup(c *gin.Context)  {}
func (u *UserHandler) Login(c *gin.Context)   {}
func (u *UserHandler) Edit(c *gin.Context)    {}
func (u *UserHandler) Profile(c *gin.Context) {}

// 路由注册：

func (u *UserHandler) RegisterRouters(server *gin.Engine) {
	server.PUT("/users/signup", u.Signup)
	server.POST("/users/login", u.Login)
	server.GET("/users/edit", u.Edit)
	server.GET("/users/profile", u.Profile)

}

// 分组路由：

func (u *UserHandler) RegisterRoutersWithGroup(ug *gin.RouterGroup) {
	ug.PUT("/users/signup", u.Signup)
	ug.POST("/users/login", u.Login)
	ug.GET("/users/edit", u.Edit)
	ug.GET("/users/profile", u.Profile)
}
