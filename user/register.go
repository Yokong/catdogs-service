package main

import (
	pb "catdogs-proto"
	configs "catdogs-service/configs/common"
	"catdogs-service/libs"
	"catdogs-service/models"
	"context"
	"crypto/md5"
	"fmt"
	"strconv"
)

func RegisterHandler(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterRsp, error) {
	has := verifyUser(in)
	if has {
		return &pb.RegisterRsp{
			Code: -1000,
			Msg: "用户已存在",
			Data: []byte(""),
		}, nil
	}
	go saveUser(in)
	return &pb.RegisterRsp{
		Code: 0,
		Msg: "success",
		Data: []byte(""),
	}, nil
}

func saveUser(in *pb.RegisterReq) {
	pwData := md5.Sum([]byte(in.Password + configs.C.PwSalt))
	pwHex := fmt.Sprintf("%x", pwData)
	session := models.Db.NewSession()
	defer session.Close()
	err := session.Begin()
	u := models.User{
		Email: in.Email,
		Password: pwHex,
	}
	err = u.Set()
	if err != nil {
		fmt.Println(err)
		session.Rollback()
		return
	}
	// 生成openid
	strId := strconv.Itoa(u.Id)
	openid := libs.AESEncrypt(strId)
	newu := models.User{Openid: openid}
	_, err = session.Id(u.Id).Update(newu)
	if err != nil {
		fmt.Println(err)
		session.Rollback()
		return
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 验证用户是否存在
func verifyUser(in *pb.RegisterReq) bool {
	u := models.User{
		Email: in.Email,
	}
	has, err := u.Get()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return has
}
