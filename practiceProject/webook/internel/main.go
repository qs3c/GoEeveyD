package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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
		// 跨域请求的“请求头”中可以有哪些字段
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// 指定哪些非简单“响应头”可以被浏览器的脚本访问
		ExposeHeaders:    []string{"x-jwt-token"},
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
	//store := cookie.NewStore([]byte("secret"))

	// session memstore 内存（注意和memcache区分，memcache是数据库）
	//store := memstore.NewStore([]byte("H823kgHYwvHm9BltzLty2ZFU0vxBPVpg"), []byte("nm6KfIptEwIDOZjXONqB0a3ImWJeR4tOY"))
	// session redis
	store, err := redis.NewStore(16, "tcp",
		"localhost:6379", "",
		[]byte("H823kgHYwvHm9BltzLty2ZFU0vxBPVpg"), []byte("nm6KfIptEwIDOZjXONqB0a3ImWJeR4tOY"))
	if err != nil {
		panic(err)
	}
	server.Use(sessions.Sessions("mysession", store))

	// 登录状态预验证（已登录情况，看看 session 有没有已经 set 过
	//server.Use(middleware.NewLoginMiddlewareBuilder().Build())
	server.Use(middleware.NewLoginJWTMiddlewareBuilder().Build())
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
