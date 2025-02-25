package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
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
type Register struct {
	Phone_number string `json:"phone_number"`
	Password     string `json:"password"`
	User_name    string `json:"user_name"`
}
type Login struct {
	Phone_number string `json:"phone_number"`
	Password     string `json:"password"`
}
type RegisterRsp struct {
	Code int32
	Msg  string
}
type LoginRsp struct {
	Code int32
	Msg  string
	User User
}

type UserRepo interface {
	Create(context.Context, *User) error
	Update(context.Context, int64, *User) error
	IsExist(context.Context, int64) (bool, error)
	Delete(context.Context, int64) error
	Find(context.Context, *FindParams) (*User, error)
	Register(context.Context, *Register) error
	Login(context.Context, *Login) (*User, error)
	IsExistbyPhone(context.Context, string) (bool, error)
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

func (uc *UserUsecase) RegisterUser(ctx context.Context, g *Register) (*RegisterRsp, error) {
	uc.log.WithContext(ctx).Infof("RegisterUser: %+v", g)
	//先判断改账号，是否已经被注册了
	ok, err := uc.repo.IsExistbyPhone(ctx, g.Phone_number)
	if err != nil {
		return &RegisterRsp{
			Code: 400,
			Msg:  "注册发生错误",
		}, err
	}
	if ok {
		return &RegisterRsp{
			Code: 400,
			Msg:  "注册失败，账号已注册",
		}, err
	}
	err = uc.repo.Register(ctx, g)
	if err != nil {
		return &RegisterRsp{
			Code: 400,
			Msg:  "注册失败，数据库错误",
		}, err
	}
	return &RegisterRsp{
		Code: 0,
		Msg:  "注册成功，请跳转登录页面",
	}, nil

}
func (uc *UserUsecase) LoginUser(ctx context.Context, g *Login) (*LoginRsp, error) {
	uc.log.WithContext(ctx).Infof("LoginUser: %+v", g)
	//先判断是否存在该用户
	ok, err := uc.repo.IsExistbyPhone(ctx, g.Phone_number)
	if err != nil {
		return &LoginRsp{
			Code: 400,
			Msg:  "登录发生错误",
			User: User{},
		}, err
	}
	if !ok {
		return &LoginRsp{
			Code: 400,
			Msg:  "登录失败，账号不存在，请先注册",
			User: User{},
		}, err
	}
	//再去查找这个用户
	users, err := uc.repo.Login(ctx, g)
	if err != nil {
		return nil, err
	}
	//再判断密码是否正确
	//如果存在用户，则比较数据库中的密码，和登录传的密码，是否一致
	if err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(g.Password)); err != nil {
		return &LoginRsp{
			Code: 400,
			Msg:  "登录失败，密码错误",
			User: User{},
		}, err
	}
	//密码也正确，登录成功，并返回用户的信息
	return &LoginRsp{
		Code: 0,
		Msg:  "登录成功",
		User: User{
			ID:           users.ID,
			Phone_number: users.Phone_number,
			User_name:    users.User_name,
			User_type:    users.User_type,
			Img_url:      users.Img_url,
			Description:  users.Description,
		},
	}, nil
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
