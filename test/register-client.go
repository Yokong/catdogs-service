package main

import (
	pb "catdogs-proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)

	r, err := c.Register(context.Background(), &pb.RegisterReq{
		Email:    "18836617@qq.com",
		Password: "123123123",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Code)
	fmt.Println(r.Msg)
	fmt.Println(r.Data)
}
