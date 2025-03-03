package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"payment_system/api/order"
	"payment_system/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrderClient, NewPaymentRepo, NewOrderStatusRepo)

// Data .用于存储支付数据
type Data struct {
	// TODO wrapped database orderClient
	db     *gorm.DB
	logger log.Logger
}

// OrderClient .用于调用订单服务
type OrderClient struct {
	client order.OrderServiceClient
	logger log.Logger
}

type PaymentInfo struct {
	OrderId uint64  `gorm:"order_id"`
	Amount  float32 `gorm:"amount"`
	Status  string  `gorm:"status"`
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// 连接数据库
	databaseAddr := os.Getenv("MYSQL_ADDR")
	if databaseAddr == "" {
		databaseAddr = c.GetDatabase().GetSource()
	}
	log.Infof("开始连接数据库 %+v", databaseAddr)

	mysqlDB, er := gorm.Open(mysql.Open(databaseAddr))
	if er != nil {
		panic(er)
	}
	if er := mysqlDB.AutoMigrate(&PaymentInfo{}); er != nil {
		panic(er)
	} //自动迁移，自动创建表，默认蛇行复数 命名
	//拿到mysqlDB的实例
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100) //最大连接数
	db.SetMaxIdleConns(50)  //最大空闲连接数，一般为最大连接数/2
	return &Data{
		db:     mysqlDB,
		logger: logger,
	}, cleanup, nil
}

func NewOrderClient(c *conf.Data, logger log.Logger) (*OrderClient, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// new etcd orderClient
	etcdAddr := os.Getenv("ETCD_ADDR")
	if etcdAddr == "" {
		etcdAddr = "127.0.0.1:2379" // 测试环境
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdAddr}, //从etcd中，服务的发现
	})
	if err != nil {
		panic(err)
	}
	// new dis with etcd orderClient
	dis := etcd.New(client)

	//endpoint := "discovery:///provider"
	endpoint := "discovery:///order_service" //把etcd的Name标识符拿过来，找到对应的服务ip
	conn, err := grpc.DialInsecure(
		context.Background(),
		//grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		//上面127.0.0.1:9000写死的，后面引入etcd服务发现。把服务注册到etcd中
		//我们的服务，现在微服务化了，有很多的分布式节点，需要etcd帮我们存储现在所有的分布式节点
		//实现负载均衡的能力，指定请求的节点，每个请求路由到不同的机器节点上
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis), //服务的发现
	)
	if err != nil {
		log.Infof("注册订单微服务失败 etcd addr:%+v。 可能由于etcd服务不可达", etcdAddr)
		panic(err)
	}
	orderClient := order.NewOrderServiceClient(conn)
	log.Infof("成功注册订单微服务。 etcd addr:%+v, connection: %+v", etcdAddr, orderClient)

	return &OrderClient{
		logger: logger,
		client: orderClient,
	}, cleanup, nil
}
