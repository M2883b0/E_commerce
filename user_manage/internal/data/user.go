package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
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
	Address      string `gorm:"column:address;"`
}

func (UserDetail) TableName() string {
	table := "ec.user" //数据库的表名
	return table
}

func (c *userRepo) Create(ctx context.Context, user *biz.User) error {
	c.log.Infof("内部用户创建 = %+v", user)
	//密码加密
	hashedPassword, err := encryptPassword(user.Password)
	if err != nil {
		c.log.Errorf("密码加密错误 = %v", err)
		return err
	}
	detail := UserDetail{
		Phone_number: user.Phone_number,
		Password:     hashedPassword,
		User_name:    user.User_name,
		User_type:    user.User_type,
		Img_url:      user.Img_url,
		Description:  user.Description,
		Address:      user.Address,
	}
	db := c.data.db
	if err := db.Create(&detail).Error; err != nil {
		c.log.Errorf("内部用户创建错误 = %v", err)
		return err
	}

	return nil
}

func (c *userRepo) Register(ctx context.Context, user *biz.Register) error {
	c.log.Infof("注册用户 = %+v", user)
	//密码加密
	hashedPassword, err := encryptPassword(user.Password)
	if err != nil {
		c.log.Errorf("密码加密错误 = %v", err)
		return err
	}
	detail := UserDetail{
		Phone_number: user.Phone_number,
		Password:     hashedPassword,
		User_name:    user.User_name,
		User_type:    1,
		Img_url:      "default.jpg",
		Description:  "Hello World",
		Address:      "",
	}
	db := c.data.db
	if err := db.Create(&detail).Error; err != nil {
		c.log.Errorf("登录错误 = %v", err)
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
		Password:     results.Password,
		User_name:    results.User_name,
		User_type:    results.User_type,
		Img_url:      results.Img_url,
		Description:  results.Description,
		Address:      results.Address,
	}
	return users, nil
}

func (c *userRepo) Update(ctx context.Context, id int64, user *biz.User) (int32, string, error) {
	c.log.Infof("用户请求更新信息 = %+v", user)
	db := c.data.db
	//密码加密
	hashedPassword, err := encryptPassword(user.Password)
	if err != nil {
		c.log.Errorf("密码加密错误 = %v", err)
		return 400, "服务器错误，密码加密失败", err
	}
	if user.Password == "" { //如果不更新密码
		hashedPassword = "" //加密后，强行置为空
	}
	if user.Phone_number != "" {
		//查看是否有这个账号
		if isExist, err := c.IsExistbyPhone(ctx, user.Phone_number); err != nil {
			c.log.Infof("用户查询错误 = %v", err)
			return 400, "服务器错误，查询失败", err
		} else if isExist {
			return 400, "更新失败，该账号已被占用", nil
		}
	}
	detail := UserDetail{
		Phone_number: user.Phone_number,
		Password:     hashedPassword,
		User_name:    user.User_name,
		User_type:    user.User_type,
		Img_url:      user.Img_url,
		Description:  user.Description,
		Address:      user.Address,
	}
	if err := db.Where("id = ?", id).Updates(&detail).Error; err != nil {
		c.log.Infof("用户更新错误 = %v", err)
		return 400, "服务器错误，用户信息更新失败", err
	}
	return 0, "更新成功", nil
}

func (c *userRepo) IsExist(ctx context.Context, id int64) (bool, error) {
	c.log.Infof("用户请求注销账号 = %+v", id)
	db := c.data.db
	var detail UserDetail
	err := db.Where("id = ?", id).First(&detail).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		c.log.Infof("用户不存在 = %v", err)
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
		c.log.Infof("用户不存在 = %v", err)
		return false, err
	}
	return true, nil
}

func (c *userRepo) Delete(ctx context.Context, id int64) (int32, string, error) {
	db := c.data.db
	// 删除索引信息
	err := db.Where("id = ?", id).
		Delete(&UserDetail{}).Error
	if err != nil {
		c.log.Infof("用户删除错误 = %v", err)
		return 400, "服务器错误，用户信息删除失败", err
	}
	return 0, "删除成功", nil
}

func (c *userRepo) Find(ctx context.Context, params *biz.FindParams) (*biz.User, int32, string, error) {
	query := c.data.db.Model(&UserDetail{})
	// 构造查询条件
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	var results *UserDetail
	if err := query.Find(&results).Error; err != nil {
		c.log.Infof("用户查询错误 = %v", params)
		return nil, 400, "服务器错误，查询失败", err
	}
	var users *biz.User
	users = &biz.User{
		ID:           int64(results.ID),
		Phone_number: results.Phone_number,
		User_name:    results.User_name,
		User_type:    results.User_type,
		Img_url:      results.Img_url,
		Description:  results.Description,
		Address:      results.Address,
	}
	return users, 0, "查询成功", nil
}

func encryptPassword(password string) (string, error) {
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if error != nil {
		log.Infof("密码加密错误: %v", error)
		return "", error
	}
	return string(hashedPassword), nil
}

//操作db
