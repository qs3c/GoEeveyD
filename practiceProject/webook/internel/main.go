package main

import (
	"github.com/gin-gonic/gin"
	"practiceProject/webook/internel/web"
)

func main() {
	server := gin.Default()
	u := &web.UserHandler{}
	u.RegisterRouters(server)
	_ = server.Run("localhost:8080")
}
