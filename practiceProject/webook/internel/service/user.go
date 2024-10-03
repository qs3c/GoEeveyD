package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"practiceProject/webook/internel/domain"
	"practiceProject/webook/internel/repository"
)

// 每一层都有自己的 ErrDuplicateEmail 用下一层做的一个别名
// 设计原则是每一层只会跟下一层交互，有比较好的隔离性
// 比全局变量的做法要好

var (
	ErrDuplicateEmail        = repository.ErrDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("账号或密码不对！")
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
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	// 2 和存储层 repository 交互进行数据存储

	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, email, password string) (domain.User, error) {

	// 通过账号从数据库中查出账号密码
	u, err := svc.repo.FindByEmail(ctx, email)

	// 无用户记录
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}

	// 比较输入的密码 password 和数据库中的密码 u.Password 是否一致
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}
