package libs

import (
	pb "catdogs-service/pb"
)

type R struct {
	Code int
	Msg  string
}

func GenRsp(r *R) *pb.Rsp {
	msg := ""
	if r.Msg != "" {
		msg = r.Msg
	} else if m, ok := Codes[r.Code]; ok {
		msg = m
	}
	rsp := pb.Rsp{
		Code: int32(r.Code),
		Msg:  msg,
	}
	return &rsp
}
