package main

import (
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"

	"content_manage/internal/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
// etcd注册的一些Key
var (
	// Name is the name of the compiled software.
	Name = "content_manage"
	// Version is the version of the compiled software.
	Version = "1.0.0"
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(c *conf.Data, logger log.Logger, gs *grpc.Server) *kratos.App {
	etcdAddr := os.Getenv("ETCD_ADDR")
	if etcdAddr == "" {
		etcdAddr = c.Redis.GetAddr()
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdAddr}, //本地的etcd服务
	})
	reg := etcd.New(client)
	if err != nil {
		panic(err)
	}
	return kratos.New(
		kratos.ID(id),     //本台系统的名字
		kratos.Name(Name), //自己起的etcd的name标识符
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(reg), //把grpc的9000端口服务，注册到etcd中
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
