package main

import (
	pb "catdogs-proto"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50001"
)

type User struct{}

func (u *User) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterRsp, error) {
	fmt.Println(in.Email)
	fmt.Println(in.Password)
	return &pb.RegisterRsp{
		Code: 0,
		Msg:  "success",
		Data: []byte(""),
	}, nil
}

func main() {
	initServer()
}

func initServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &User{})
	s.Serve(lis)
}
