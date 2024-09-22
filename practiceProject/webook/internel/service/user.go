package service

import (
	"context"
	"practiceProject/webook/internel/domain"
	"practiceProject/webook/internel/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

// service 初始化函数

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// 服务层的功能函数，后续会给 web 层的 handler 去调用

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	// 注册服务做什么
	// 1 加密
	// 2 和存储层 repository 交互进行数据存储
	return nil
}
