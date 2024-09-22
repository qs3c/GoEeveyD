package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type UserDAO struct {
	// 最后一层，直接和数据库打交道的层
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

// 封装的一个操作 gorm 的函数

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	// 毫秒数
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	// 一个 gorm 的 Create 操作
	return dao.db.WithContext(ctx).Create(&u).Error

}

// 对标到数据库的 User 结构体（和 domain 中的有什么区别？domain 中的是业务 User）
// 直接对应于数据库表

type User struct {
	Id int64 `gorm:"primary_key,auto_increment"`
	// 邮箱是唯一的，所以用唯一索引
	Email    string `gorm:"type:varchar(255);unique"`
	Password string

	// 创建时间 毫秒数
	Ctime int64
	// 更新时间 毫秒数
	Utime int64
}
