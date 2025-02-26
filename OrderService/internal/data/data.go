package data

import (
	"OrderService/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrderRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	mysqlDB, er := gorm.Open(mysql.Open(c.GetDatabase().GetSource()))
	if er != nil {
		panic(er)
	}
	if er := mysqlDB.AutoMigrate(&OrderInfo{}); er != nil {
		panic(er)
	} //自动迁移，自动创建表，默认蛇行复数 命名
	//拿到mysqlDB的实例
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100) //最大连接数
	db.SetMaxIdleConns(50)  //最大空闲连接数，一般为最大连接数/2
	return &Data{db: mysqlDB}, cleanup, nil
}
