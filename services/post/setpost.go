package main

import (
	"catdogs-service/libs"
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
		rsp.Rsp = libs.GenRsp(&libs.R{Code: -999})
		return err
	}

	rsp.Rsp = libs.GenRsp(&libs.R{Code: 0})
	return nil
}
