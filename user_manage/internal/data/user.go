package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"user_manage/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type UserDetail struct {
	gorm.Model
	Phone_number string `gorm:"column:phone_number;"`
	Password     string `gorm:"column:password;"`
	User_name    string `gorm:"column:user_name;"`
	User_type    int32  `gorm:"column:user_type;"`
	Img_url      string `gorm:"column:img_url;"`
	Description  string `gorm:"column:description;"`
}

func (UserDetail) TableName() string {
	table := "ec.user" //数据库的表名
	return table
}

func (c *userRepo) Create(ctx context.Context, user *biz.User) error {
	c.log.Infof("userRepo Create user = %+v", user)
	detail := UserDetail{
		Phone_number: user.Phone_number,
		Password:     user.Password,
		User_name:    user.User_name,
		User_type:    user.User_type,
		Img_url:      user.Img_url,
		Description:  user.Description,
	}
	db := c.data.db
	if err := db.Create(&detail).Error; err != nil {
		c.log.Errorf("user create error = %v", err)
		return err
	}

	return nil
}

func (c *userRepo) Register(ctx context.Context, user *biz.Register) error {
	c.log.Infof("userRepo Register user = %+v", user)
	detail := UserDetail{
		Phone_number: user.Phone_number,
		Password:     user.Password,
		User_name:    user.User_name,
	}
	db := c.data.db
	if err := db.Create(&detail).Error; err != nil {
		c.log.WithContext(ctx).Errorf("user create error = %v", err)
		return err
	}
	return nil
}

func (c *userRepo) Login(ctx context.Context, user *biz.Login) (*biz.User, error) {
	c.log.Infof("userRepo Login user = %+v", user)
	db := c.data.db
	//查询用户信息
	var results UserDetail
	if err := db.Where("phone_number = ?", user.Phone_number).Find(&results).Error; err != nil {
		c.log.WithContext(ctx).Errorf("user login error = %v", err)
		return nil, err
	}
	//返回查询的用户信息
	var users *biz.User
	users = &biz.User{
		ID:           int64(results.ID),
		Phone_number: results.Phone_number,
		User_name:    results.User_name,
		User_type:    results.User_type,
		Img_url:      results.Img_url,
		Description:  results.Description,
	}
	return users, nil
}

func (c *userRepo) Update(ctx context.Context, id int64, user *biz.User) error {
	c.log.Infof("userRepo Update user = %+v", user)
	detail := UserDetail{
		Phone_number: user.Phone_number,
		Password:     user.Password,
		User_name:    user.User_name,
		User_type:    user.User_type,
		Img_url:      user.Img_url,
		Description:  user.Description,
	}
	db := c.data.db
	if err := db.Where("id = ?", id).Updates(&detail).Error; err != nil {
		c.log.WithContext(ctx).Errorf("user update error = %v", err)
		return err
	}
	return nil
}

func (c *userRepo) IsExist(ctx context.Context, id int64) (bool, error) {
	db := c.data.db
	var detail UserDetail
	err := db.Where("id = ?", id).First(&detail).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		c.log.WithContext(ctx).Errorf("User isExist = [%v]", err)
		return false, err
	}
	return true, nil
}

func (c *userRepo) IsExistbyPhone(ctx context.Context, phone_number string) (bool, error) {
	db := c.data.db
	var detail UserDetail
	err := db.Where("phone_number = ?", phone_number).First(&detail).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		c.log.WithContext(ctx).Errorf("User isExist = [%v]", err)
		return false, err
	}
	return true, nil
}

func (c *userRepo) Delete(ctx context.Context, id int64) error {
	db := c.data.db
	// 删除索引信息
	err := db.Where("id = ?", id).
		Delete(&UserDetail{}).Error
	if err != nil {
		c.log.WithContext(ctx).Errorf("user delete error = %v", err)
		return err
	}
	return nil
}

func (c *userRepo) Find(ctx context.Context, params *biz.FindParams) (*biz.User, error) {
	query := c.data.db.Model(&UserDetail{})
	// 构造查询条件
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	var results *UserDetail
	if err := query.Find(&results).Error; err != nil {
		c.log.WithContext(ctx).Errorf("user find error = %v", err)
		return nil, err
	}
	var users *biz.User
	users = &biz.User{
		ID:           int64(results.ID),
		Phone_number: results.Phone_number,
		User_name:    results.User_name,
		User_type:    results.User_type,
		Img_url:      results.Img_url,
		Description:  results.Description,
	}
	return users, nil
}

//操作db
