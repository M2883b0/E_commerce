package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"user_manage/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	if mysqlAddr == "" {
		mysqlAddr = c.GetDatabase().GetSource()
	}
	mysqlDB, er := gorm.Open(mysql.Open(mysqlAddr))
	if er != nil {
		panic(er)
	}
	err := mysqlDB.AutoMigrate(&UserDetail{})
	if err != nil {
		return nil, nil, err
	} //自动迁移，自动创建表，默认蛇行负数
	//拿到mysqlDB的实例
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100) //最大连接数
	db.SetMaxIdleConns(50)  //最大空闲连接数，一般为最大连接数/2
	return &Data{db: mysqlDB}, cleanup, nil
}
