package data

import (
	"checkout_system/api/operate"
	"checkout_system/internal/conf"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
)

// ProviderSet is data providers.
//var ProviderSet = wire.NewSet(NewData, NewProductClient, NewCheckoutRepo)

var ProviderSet = wire.NewSet(NewProductClient, NewCheckoutRepo)

//// Data .
//type Data struct {
//	// TODO wrapped database client
//	log *log.Helper
//	//productClient operate.AppClient
//}

// // NewData .
//func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
//	cleanup := func() {
//		log.NewHelper(logger).Info("closing the data resources")
//	}
//
//	return &Data{}, cleanup, nil
//}

type ProductClient struct {
	log    *log.Helper
	client operate.AppClient
}

func NewProductClient(c *conf.Data, logger log.Logger) (*ProductClient, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// new etcd client
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
	// new dis with etcd client
	dis := etcd.New(client)

	//endpoint := "discovery:///provider"
	endpoint := "discovery:///content_manage" //把etcd的Name标识符拿过来，找到对应的服务ip
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
		log.Infof("注册商品微服务失败 etcd addr:%+v。 可能由于etcd服务不可达", etcdAddr)
		panic(err)
	}
	productClient := operate.NewAppClient(conn)

	log.Infof("成功注册商品微服务。 etcd addr:%+v, connection: %+v", etcdAddr, productClient)

	return &ProductClient{
		log:    log.NewHelper(logger),
		client: productClient,
	}, cleanup, nil
}
