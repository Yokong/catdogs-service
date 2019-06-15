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
	"strconv"
)

func RegisterHandler(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterRsp, error) {
	logging.Info(in.Email, " commit register")
	has, err := verifyUser(in)
	if err != nil {
		return &pb.RegisterRsp{
			Code:  -999,
			Msg:   "服务器出了点问题",
			Token: "",
		}, nil
	}
	if has {
		return &pb.RegisterRsp{
			Code:  -1000,
			Msg:   "用户已存在",
			Token: "",
		}, nil
	}

	tokenCh := make(chan string)
	go saveUser(in, tokenCh)
	token := <-tokenCh

	logging.Info(in.Email, " done register")
	return &pb.RegisterRsp{
		Code:  0,
		Msg:   "success",
		Token: token,
	}, nil
}

func saveUser(in *pb.RegisterReq, tokenCh chan string) {
	pwData := md5.Sum([]byte(in.Password + configs.C.PwSalt))
	pwHex := fmt.Sprintf("%x", pwData)
	session := models.Db.NewSession()
	defer session.Close()
	err := session.Begin()
	u := models.User{
		Email:    in.Email,
		Password: pwHex,
	}
	err = u.Set()
	if err != nil {
		logging.Error("User Set: ", err)
		session.Rollback()
		return
	}
	// 生成openid
	strId := strconv.Itoa(u.Id)
	openid := libs.AESEncrypt(strId)
	newu := models.User{Openid: openid}
	_, err = session.Id(u.Id).Update(newu)
	if err != nil {
		logging.Error("Update New User: ", err)
		session.Rollback()
		return
	}
	err = session.Commit()
	if err != nil {
		logging.Error("Session Commit: ", err)
		return
	}

	// 生成token
	token, err := libs.GenerateToken(openid)
	if err != nil {
		logging.Error("Generate Token: ", err)
	}
	tokenCh <- token
}

// 验证用户是否存在
func verifyUser(in *pb.RegisterReq) (bool, error) {
	u := models.User{
		Email: in.Email,
	}
	has, err := u.Get()
	if err != nil {
		logging.Error("VerifyUser Get: ", err)
		return false, err
	}
	return has, nil
}
