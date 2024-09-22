package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"practiceProject/webook/internel/repository"
	"practiceProject/webook/internel/repository/dao"
	"practiceProject/webook/internel/service"
	"practiceProject/webook/internel/web"
)

func main() {

	// 初始化过程
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
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)

	server := gin.Default()
	//u := &web.UserHandler{}

	u.RegisterRouters(server)

	_ = server.Run("localhost:8080")
}
