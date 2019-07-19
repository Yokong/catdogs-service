package main

import (
	pb "catdogs-proto"
	"catdogs-service/models"
	"context"
	"fmt"
	"time"
)

func PosterHandler(ctx context.Context, req *pb.SetPostReq, rsp *pb.SetPostRsp) error {
	rsp.Code = 0
	rsp.Msg = "success"
	fmt.Println(req.Content)
	post := models.Post{
		Title:     req.Title,
		Content:   []byte(req.Content),
		Author:    req.Author,
		Timestamp: int(time.Now().Unix()),
	}
	err := post.Set()
	if err != nil {
		return err
	}
	return nil
}
