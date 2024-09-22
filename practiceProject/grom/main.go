package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	db, err = gorm.Open(mysql.Open("root:root@tcp(localhost:13306)/xiaohongshu"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 开启 debug 模式
	db = db.Debug()

	// 迁移 schema
	_ = db.AutoMigrate(&Product{})

	// 创建一条数据
	db.Create(&Product{Code: "d42", Price: 100})
	// 查找第一个符合条件的数据，并把查到的结果放到这个 product 实例中
	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "d42")

	// 更新 price 为 200
	db.Model(&product).Update("Price", 300)
	// 更新多个字段用 Updates
	db.Model(&product).Updates(Product{Price: 200, Code: "f23"})

	// 删除
	db.Delete(&product, 1)

}
