package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"practiceProject/webook/internel/repository"
	"practiceProject/webook/internel/repository/dao"
	"practiceProject/webook/internel/service"
	"practiceProject/webook/internel/web"
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
