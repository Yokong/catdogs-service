package main

import (
	pb "catdogs-proto"
	"catdogs-service/models"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	addr = ":50001"
)

type User struct{}

func (u *User) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterRsp, error) {
	return RegisterHandler(ctx, in)
}

func (u *User) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRsp, error) {
	return LoginHandler(ctx, in)
}

func init() {
	models.InitModel()
}

func main() {
	initServer()
}

func initServer() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &User{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("failed to Serve: ", err)
	}
}
