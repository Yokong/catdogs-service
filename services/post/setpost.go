package main

import (
	pb "catdogs-proto"
	"catdogs-service/logging"
	"catdogs-service/models"
	"context"
	"time"
)

func PosterHandler(ctx context.Context, req *pb.SetPostReq, rsp *pb.SetPostRsq) error {
	post := models.Post{
		Title:     req.Title,
		Content:   req.Content,
		Author:    req.Author,
		Timestamp: time.Now().Unix(),
	}
	err := post.Set()
	if err != nil {
		logging.Error("SET POST: ", err)
	}
}
