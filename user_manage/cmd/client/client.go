package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"log"
	"user_manage/api/operate"
)

func main() {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:59815"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := operate.NewUserClient(conn)

	//增
	//reply, err := client.CreateUser(context.Background(), &operate.CreateUserRequest{
	//	User: &operate.UserInfo{
	//		PhoneNumber: "123",
	//		Password:    "123",
	//		UserName:    "lyx",
	//		UserType:    0,
	//		ImgUrl:      "xxx.jpg",
	//		Description: "111",
	//	},
	//})
	//改
	//reply, err := client.UpdateUser(context.Background(), &operate.UpdateUserRequest{
	//	User: &operate.UserInfo{
	//		Id:          1,
	//		PhoneNumber: "191",
	//		Password:    "",
	//		UserName:    "",
	//		UserType:    0,
	//		ImgUrl:      "",
	//		Description: "",
	//	},
	//})
	//删
	reply, err := client.DeleteUser(context.Background(), &operate.DeleteUserRequest{
		Id: 1,
	})
	//查
	//reply, err := client.GetUser(context.Background(), &operate.GetUserRequest{
	//	Id: 1,
	//})

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] user %+v\n", reply)
}
