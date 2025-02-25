package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID           int64  `json:"id"`
	Phone_number string `json:"phone_number"`
	Password     string `json:"password"`
	User_name    string `json:"user_name"`
	User_type    int32  `json:"user_type"`
	Img_url      string `json:"img_url"`
	Description  string `json:"description"`
}

type UserRepo interface {
	Create(context.Context, *User) error
	Update(context.Context, int64, *User) error
	IsExist(context.Context, int64) (bool, error)
	Delete(context.Context, int64) error
	Find(context.Context, *FindParams) (*User, error)
}

// 查找的参数
type FindParams struct {
	ID int64
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) error {
	uc.log.WithContext(ctx).Infof("CreateUser: %+v", g)
	return uc.repo.Create(ctx, g)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, g *User) error {
	uc.log.WithContext(ctx).Infof("UpdateUser: %+v", g)
	return uc.repo.Update(ctx, int64(g.ID), g)
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id int64) error {
	uc.log.WithContext(ctx).Infof("DeleteUser: %+v", id)
	//复合操作，现判断该用户是否存在，再执行删除操作
	ok, err := uc.repo.IsExist(ctx, id)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("User no exist")
	}
	//用户存在的情况,执行删除操作
	return uc.repo.Delete(ctx, id)
}

func (uc *UserUsecase) FindUser(ctx context.Context, params *FindParams) (*User, error) {
	users, err := uc.repo.Find(ctx, params)
	if err != nil {
		return nil, err
	}
	return users, nil
}

//执行组合逻辑
