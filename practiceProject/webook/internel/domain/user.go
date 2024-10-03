package domain

import "time"

// 用户对象信息（在 domain 中定义，然后在 service 等地方使用，不在 service 中定义

type User struct {
	Id       int64
	Email    string
	Password string
	Ctime    time.Time
}
