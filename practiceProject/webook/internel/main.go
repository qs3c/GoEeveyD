package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"practiceProject/webook/internel/repository"
	"practiceProject/webook/internel/repository/dao"
	"practiceProject/webook/internel/service"
	"practiceProject/webook/internel/web"
	"practiceProject/webook/internel/web/middleware"
	"strings"
	"time"
)

func main() {

	// 初始化过程
	// 初始化数据库并建表
	db := initDB()
	// 初始化服务器并  Use 中间件
	server := initWebServer()

	u := initUser(db)
	// 注册路由
	u.RegisterRouters(server)
	_ = server.Run("localhost:8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	//u := &web.UserHandler{}
	// 跨域中间件
	server.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "youcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	// session cookie 中间件
	// 创建一个叫 secret 的 cookie 用来存 session 信息
	// 创建了一个叫 myseesion 用作于 seesion 的 cookie
	// 后续 cookie 的 key 值是  mysession
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", store))

	// 登录状态预验证（已登录情况，看看 session 有没有已经 set 过
	server.Use(middleware.NewLoginMiddlewareBuilder().Build())
	// 注册路由
	//u.RegisterRouters(server)
	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		// panic 意味着 goroutine 直接结束
		// 只会在初始化过程中出错的时候用 panic
		// 一旦初始化出错，应用就不要启动了
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
