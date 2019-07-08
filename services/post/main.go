package main

import (
	pb "catdogs-proto"
	"context"
	"fmt"

	"github.com/micro/go-micro"
)

type Post struct{}

func (p *Post) Poster(ctx context.Context, req *pb.SetPostReq, rsp *pb.SetPostRsp) error {
	rsp.Code = 0
	rsp.Msg = "Success"
	return nil
}

func main() {
	initServer()
}

func initServer() {
	service := micro.NewService(micro.Name("post"))
	service.Init()

	err := pb.RegisterPostHandler(service.Server(), new(Post))
	if err != nil {
		fmt.Println(err)
	}
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
