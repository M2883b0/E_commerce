package main

import (
	"auth_manage/api/operate"
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"log"
)

func main() {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := operate.NewAuthClient(conn)

	//加密
	//reply, err := client.DeliverTokenByRPC(context.Background(), &operate.DeliverTokenReq{
	//	UserId: "lyx123",
	//})
	//验证
	reply, err := client.VerifyTokenByRPC(context.Background(), &operate.VerifyTokenReq{
		Token: "",
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] user %+v\n", reply)
}
