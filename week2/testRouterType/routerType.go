package testRouterType

import "github.com/gin-gonic/gin"

func RouterType() {
	server := gin.Default()
	// 1. 静态路由
	server.GET("/static/get", func(c *gin.Context) {
		c.String(200, "get.html")
	})
	server.POST("/static/post", func(c *gin.Context) {
		c.String(200, "post.html")
	})

	// 2. 参数路由
	server.DELETE("/param/delete/:name", func(c *gin.Context) {
		// 用 Param 来获得
		name := c.Param("name")
		c.String(200, "delete.html"+name)

	})

	// 3. 通配符路由
	server.PUT("/param/put/*.html", func(c *gin.Context) {
		// 用 Param 来获得
		page := c.Param(".html")
		c.String(200, "put.html"+page)
	})

	// 新增一个查询参数 Query
	server.HEAD("/head", func(c *gin.Context) {
		oid := c.Query("id")
		c.String(200, "head.html"+oid)
	})

	_ = server.Run("localhost:8080")
}
