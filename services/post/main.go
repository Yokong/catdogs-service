package main

import (
	"catdogs-service/models"
	pb "catdogs-service/pb"
	"context"
	"fmt"

	"github.com/micro/go-micro"
)

type Post struct{}

func (p *Post) Poster(ctx context.Context, req *pb.SetPostReq, rsp *pb.SetPostRsp) error {
	return PosterHandler(ctx, req, rsp)
}

func (p *Post) GetPostById(ctx context.Context, req *pb.GetPostByIdReq, rsp *pb.GetPostByIdRsp) error {
	return nil
}

func init() {
	models.InitModel()
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
