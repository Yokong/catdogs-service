package main

import (
	"catdogs-service/logging"
	"catdogs-service/models"
	pb "catdogs-service/pb"
	"context"
	"time"
)

func PosterHandler(ctx context.Context, req *pb.SetPostReq, rsp *pb.SetPostRsp) error {
	post := models.Post{
		Title:     req.Title,
		Content:   []byte(req.Content),
		Author:    req.Author,
		Timestamp: int(time.Now().Unix()),
	}
	err := post.Set()
	if err != nil {
		logging.Error("Set Post: ", err)
		rsp.Code = -999
		rsp.Msg = "服务器出现问题"
		return err
	}

	rsp.Code = 0
	rsp.Msg = "success"
	return nil
}
