package api

import (
	"Gateway/internal/api/operate"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"net/http"
)

const (
	SessionKey = "session_id" //鉴权的id，前端需要在Headers中加入这个参数session_id
)

type JWT struct {
	operateAuthClient operate.AuthClient
}

func NewJWTAuth() *JWT {
	jwt := &JWT{}              //创建结构体实例
	connOperateAuthClient(jwt) //给实例加上鉴权grpc服务
	return jwt
}

func (s *JWT) Auth(c *gin.Context) {
	sessionID := c.GetHeader(SessionKey)
	rsp, err := s.operateAuthClient.VerifyTokenByRPC(c, &operate.VerifyTokenReq{Token: sessionID})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "jwt 鉴权失败")
		return
	}
	if !rsp.Res {
		c.AbortWithStatusJSON(http.StatusUnauthorized, rsp.Msg)
	}
	c.Set("user_id", rsp.GetUserId())
	c.Next()
}

func connOperateAuthClient(app *JWT) {
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

//// jwt中间件,只验证jwt。(生成jwt放到login_handle的登录里面)
//func JwtToken() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokenHeader := c.GetHeader(SessionKey) //从请求头，拿到token
//		if tokenHeader == "" {
//			c.AbortWithStatusJSON(http.StatusForbidden, "jwt 为空")
//			return
//		}
//		//如果不为空，则进行验证
//		key, err := jwt.CheckToken(tokenHeader)
//		if err != nil {
//			c.AbortWithStatusJSON(http.StatusInternalServerError, "jwt 鉴权失败")
//			return
//		}
//		//信息上，鉴权成功
//		//还需要判断token是否过期
//		if time.Now().Unix() > key.ExpiresAt.Unix() {
//			c.AbortWithStatusJSON(http.StatusUnauthorized, "jwt 鉴权过期")
//			return
//		}
//
//		//c.Set("username", key.Username)
//		c.Next()
//	}
//}
