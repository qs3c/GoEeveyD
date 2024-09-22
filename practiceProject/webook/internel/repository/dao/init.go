package dao

import "gorm.io/gorm"

// 初始化建表

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
