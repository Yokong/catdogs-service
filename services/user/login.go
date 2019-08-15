package main

import (
	configs "catdogs-service/configs/common"
	"catdogs-service/libs"
	"catdogs-service/logging"
	"catdogs-service/models"
	pb "catdogs-service/pb"
	"context"
	"crypto/md5"
	"fmt"
)

func LoginHandler(ctx context.Context, req *pb.LoginReq, rsp *pb.LoginRsp) error {
	logging.Info(req.Email, " entering logreq")
	u := models.User{
		Email: req.Email,
	}
	has, err := u.Get()
	if err != nil {
		logging.Error("Logreq Get User: ", err)
		rsp.Rsp = libs.GenRsp(&libs.R{Code: -999})
		rsp.Token = ""
		return nil
	}
	if !has {
		rsp.Rsp = libs.GenRsp(&libs.R{Code: -1002})
		rsp.Token = ""
		return nil
	}
	pwd := md5.Sum([]byte(req.Password + configs.C.PwSalt))
	pwdHex := fmt.Sprintf("%x", pwd)
	if pwdHex != u.Password {
		rsp.Rsp = libs.GenRsp(&libs.R{Code: -1003})
		rsp.Token = ""
		return nil
	}
	token, err := libs.GenerateToken(u.Openid)
	if err != nil {
		logging.Error("Generate Token In Logreq: ", err)
		rsp.Rsp = libs.GenRsp(&libs.R{Code: -999})
		rsp.Token = ""
		return nil
	}
	logging.Info(req.Email, " all done logreq")
	rsp.Rsp = libs.GenRsp(&libs.R{Code: 0})
	rsp.Token = token
	return nil
}
