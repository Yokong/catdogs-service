package main

import (
	pb "catdogs-proto"
	configs "catdogs-service/configs/common"
	"catdogs-service/libs"
	"catdogs-service/logging"
	"catdogs-service/models"
	"context"
	"crypto/md5"
	"fmt"
)

func LoginHandler(ctx context.Context, in *pb.LoginReq) (*pb.LoginRsp, error) {
	u := models.User{
		Email: in.Email,
	}
	has, err := u.Get()
	if err != nil {
		logging.Error("Login Get User: ", err)
		return &pb.LoginRsp{
			Code:  -999,
			Msg:   "服务器出现问题",
			Token: "",
		}, nil
	}
	if !has {
		return &pb.LoginRsp{
			Code:  -1002,
			Msg:   "用户不存在",
			Token: "",
		}, nil
	}
	pwd := md5.Sum([]byte(in.Password + configs.C.PwSalt))
	pwdHex := fmt.Sprintf("%x", pwd)
	if pwdHex != u.Password {
		return &pb.LoginRsp{
			Code:  -1003,
			Msg:   "密码错误",
			Token: "",
		}, nil
	}
	token, err := libs.GenerateToken(u.Openid)
	if err != nil {
		logging.Error("Generate Token In Login: ", err)
		return &pb.LoginRsp{
			Code:  -999,
			Msg:   "服务器出现问题",
			Token: "",
		}, nil
	}
	return &pb.LoginRsp{
		Code:  0,
		Msg:   "success",
		Token: token,
	}, nil
}
