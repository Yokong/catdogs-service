package main

import (
	pb "catdogs-proto"
	"catdogs-service/models"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

type User struct{}

func (u *User) Register(ctx context.Context, req *pb.RegisterReq, rsp *pb.RegisterRsp) error {
	rsp.Code = 0
	rsp.Msg = "Success"
	return nil
}

func init() {
	models.InitModel()
}

func main() {
	initServer()
}

func initServer() {
	service := micro.NewService(micro.Name("user"))
	service.Init()

	err := pb.RegisterUserHandler(service.Server(), new(User))
	if err != nil {
		fmt.Println(err)
	}
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
