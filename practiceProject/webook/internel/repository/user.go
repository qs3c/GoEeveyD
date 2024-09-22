package repository

import (
	"context"
	"practiceProject/webook/internel/domain"
	"practiceProject/webook/internel/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
		// Ctime Utime 属于数据库的数据 所以交给dao的层面去处理
	})
}
func (r *UserRepository) FindById(id int64) {
	// 1. cache 中寻找
	// 2. dao 中寻找
	// 3. 在 dao 中找到了之后回写 cache
}
