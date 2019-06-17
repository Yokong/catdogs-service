package main

import (
	pb "catdogs-proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

var addr = "118.24.146.34:50001"

func main() {
	// ferver := make(chan bool)
	// for i := 0; i < 1000; i++ {
	// 	login()
	// }
	// <-ferver
	// register()
	login()
}

func login() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)

	r, err := c.Login(context.Background(), &pb.LoginReq{
		Email:    "18836617@qq.com",
		Password: "513520",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Code, r.Msg)
}

func register() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)

	r, err := c.Register(context.Background(), &pb.RegisterReq{
		Email:    "153367234@qq.com",
		Password: "123123123",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Code, r.Msg)
}
