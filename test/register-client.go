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
	fmt.Println(r.Code)
	fmt.Println(r.Msg)
	fmt.Println(r.Token)
}
