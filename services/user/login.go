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

func LoginHandler(ctx context.Context, req *pb.RegisterReq, rsp *pb.RegisterRsp) error {
	logging.Info(req.Email, " enterreqg logreq")
	u := models.User{
		Email: req.Email,
	}
	has, err := u.Get()
	if err != nil {
		logging.Error("Logreq Get User: ", err)
		rsp.Code = -999
		rsp.Msg = "服务器出现问题"
		rsp.Token = ""
		return nil
	}
	if !has {
		rsp.Code = -1002
		rsp.Msg = "用户不存在"
		rsp.Token = ""
		return nil
	}
	pwd := md5.Sum([]byte(req.Password + configs.C.PwSalt))
	pwdHex := fmt.Sprintf("%x", pwd)
	if pwdHex != u.Password {
		rsp.Code = -1003
		rsp.Msg = "密码错误"
		rsp.Token = ""
		return nil
	}
	token, err := libs.GenerateToken(u.Openid)
	if err != nil {
		logging.Error("Generate Token In Logreq: ", err)
		rsp.Code = -999
		rsp.Msg = "服务器出现问题"
		rsp.Token = ""
		return nil
	}
	logging.Info(req.Email, " all done logreq")
	rsp.Code = 0
	rsp.Msg = "success"
	rsp.Token = token
	return nil
}
