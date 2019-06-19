package main

import (
	pb "catdogs-proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

//var addr = "118.24.146.34:50001"

func main() {
	// ferver := make(chan bool)
	// for i := 0; i < 1000; i++ {
	// 	login()
	// }
	// <-ferver
	register()
	//login()
}

//func login() {
//	conn, err := grpc.Dial(addr, grpc.WithInsecure())
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer conn.Close()
//
//	c := pb.NewUserClient(conn)
//
//	r, err := c.Login(context.Background(), &pb.LoginReq{
//		Email:    "18836617@qq.com",
//		Password: "513520",
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(r.Code, r.Msg)
//}

func register() {
	service := micro.NewService(micro.Name("user.client"))
	service.Init()
	user := pb.NewUserService("user", service.Client())
	rsp, err := user.Register(context.TODO(), &pb.RegisterReq{
		Email: "188",
		Password: "12312312",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Msg, rsp.Code)
}
