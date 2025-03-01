package services

import (
	"Gateway/internal/api/cart"
	"Gateway/internal/api/operate"
	"Gateway/internal/api/order"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type CmsAPP struct {
	operateAppClient   operate.AppClient
	operateUserClient  operate.UserClient
	operateAuthClient  operate.AuthClient
	orderServiceClient order.OrderServiceClient
	cartServiceClient  cart.CartServiceClient
}

func NewCmsApp() *CmsAPP {
	app := &CmsAPP{}           //创建结构体实例
	connOperateAppClient(app)  //给实例加上内容grpc服务
	connOperateUserClient(app) //给实例加上用户grpc服务
	connOperateAuthClient(app) //给实例加上鉴权grpc服务
	connOrderServiceClient(app)
	connCartServiceClient(app)
	return app
}

func connOperateAppClient(app *CmsAPP) {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"etcd-server:2379"}, //从etcd中，服务的发现
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
		panic(err)
	}
	appclient := operate.NewAppClient(conn)
	app.operateAppClient = appclient
}

func connOperateUserClient(app *CmsAPP) {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"etcd-server:2379"}, //从etcd中，服务的发现
	})
	if err != nil {
		panic(err)
	}
	// new dis with etcd client
	dis := etcd.New(client)

	//endpoint := "discovery:///provider"
	endpoint := "discovery:///user_manage" //把etcd的Name标识符拿过来，找到对应的服务ip
	conn, err := grpc.DialInsecure(
		context.Background(),
		//grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		//上面127.0.0.1:9000写死的，后面引入etcd服务发现。把服务注册到etcd中,端口自动分配
		//我们的服务，现在微服务化了，有很多的分布式节点，需要etcd帮我们存储现在所有的分布式节点
		//实现负载均衡的能力，指定请求的节点，每个请求路由到不同的机器节点上
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis), //服务的发现
	)
	if err != nil {
		panic(err)
	}
	appclient := operate.NewUserClient(conn)
	app.operateUserClient = appclient
}

func connOperateAuthClient(app *CmsAPP) {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"etcd-server:2379"}, //从etcd中，服务的发现
	})
	if err != nil {
		panic(err)
	}
	// new dis with etcd client
	dis := etcd.New(client)

	//endpoint := "discovery:///provider"
	endpoint := "discovery:///auth_manage" //把etcd的Name标识符拿过来，找到对应的服务ip
	conn, err := grpc.DialInsecure(
		context.Background(),
		//grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		//上面127.0.0.1:9000写死的，后面引入etcd服务发现。把服务注册到etcd中,端口自动分配
		//我们的服务，现在微服务化了，有很多的分布式节点，需要etcd帮我们存储现在所有的分布式节点
		//实现负载均衡的能力，指定请求的节点，每个请求路由到不同的机器节点上
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis), //服务的发现
	)
	if err != nil {
		panic(err)
	}
	appclient := operate.NewAuthClient(conn)
	app.operateAuthClient = appclient
}

func connOrderServiceClient(app *CmsAPP) {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"etcd-server:2379"}, //从etcd中，服务的发现
	})
	if err != nil {
		panic(err)
	}
	// new dis with etcd client
	dis := etcd.New(client)

	//endpoint := "discovery:///provider"
	endpoint := "discovery:///order_service" //把etcd的Name标识符拿过来，找到对应的服务ip
	conn, err := grpc.DialInsecure(
		context.Background(),
		//grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		//上面127.0.0.1:9000写死的，后面引入etcd服务发现。把服务注册到etcd中,端口自动分配
		//我们的服务，现在微服务化了，有很多的分布式节点，需要etcd帮我们存储现在所有的分布式节点
		//实现负载均衡的能力，指定请求的节点，每个请求路由到不同的机器节点上
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis), //服务的发现
	)
	if err != nil {
		panic(err)
	}
	appClient := order.NewOrderServiceClient(conn)
	app.orderServiceClient = appClient
}

func connCartServiceClient(app *CmsAPP) {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"etcd-server:2379"}, //从etcd中，服务的发现
	})
	if err != nil {
		panic(err)
	}
	// new dis with etcd client
	dis := etcd.New(client)

	//endpoint := "discovery:///provider"
	endpoint := "discovery:///cart_service" //把etcd的Name标识符拿过来，找到对应的服务ip
	conn, err := grpc.DialInsecure(
		context.Background(),
		//grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		//上面127.0.0.1:9000写死的，后面引入etcd服务发现。把服务注册到etcd中,端口自动分配
		//我们的服务，现在微服务化了，有很多的分布式节点，需要etcd帮我们存储现在所有的分布式节点
		//实现负载均衡的能力，指定请求的节点，每个请求路由到不同的机器节点上
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis), //服务的发现
	)
	if err != nil {
		panic(err)
	}
	appClient := cart.NewCartServiceClient(conn)
	app.cartServiceClient = appClient
}
