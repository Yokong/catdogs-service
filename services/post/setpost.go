package main

import (
	pb "catdogs-proto"
	"context"
	"fmt"
)

func PosterHandler(ctx context.Context, req *pb.SetPostReq, rsp *pb.SetPostRsp) error {
	rsp.Code = 0
	rsp.Msg = "success"
	fmt.Println(req.Content)
	return nil
}
