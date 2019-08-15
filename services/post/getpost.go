package main

import (
	"catdogs-service/libs"
	"catdogs-service/logging"
	"catdogs-service/models"
	pb "catdogs-service/pb"
	"context"
)

func GetPostByIdHandler(ctx context.Context, req *pb.GetPostByIdReq, rsp *pb.GetPostByIdRsp) error {
	post := models.Post{
		Id: int(req.Id),
	}
	has, err := post.Get()
	if err != nil {
		logging.Error("GetPostById Err: ", err)
		rsp.Rsp = libs.GenRsp(&libs.R{Code: -999})
		return err
	}
	if !has {
		rsp.Rsp = libs.GenRsp(&libs.R{Code: -1007})
		return nil
	}
	rsp.Rsp = libs.GenRsp(&libs.R{Code: 0})
	rsp.Author = post.Author
	rsp.Content = string(post.Content)
	rsp.Source = post.Source
	rsp.Timestamp = int64(post.Timestamp)
	return nil
}
