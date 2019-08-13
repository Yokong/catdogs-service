package main

import (
	"catdogs-service/models"
	pb "catdogs-service/pb"
	"context"
)

func SetProfileHandler(ctx context.Context, req *pb.SetProfileReq, rsp *pb.SetProfileRsp) error {
	profile := req.Profile
	p := models.Profile{
		Name:     profile.Name,
		Gender:   profile.Gender,
		Age:      profile.Age,
		PhoneNum: profile.PhoneNum,
		Email:    profile.Email,
		Birthday: profile.Birthday,
		City:     profile.City,
		Address:  profile.Address,
	}
	p.Set()

	rsp.Code = 0
	rsp.Msg = "Success"
	return nil
}
