package services

import (
	context "context"
	"fmt"
	"strconv"
)

type profileServer struct {
}

func NewProfileServer() ProfileServer {
	return profileServer{}
}

func (profileServer) mustEmbedUnimplementedProfileServer() {}

func (profileServer) GetProfile(ctx context.Context, in *GetProfileRequest) (*GetProfileResponse, error) {
	id, _ := strconv.ParseInt(in.UserId, 10, 64)
	result := fmt.Sprintf("User Id %v", id)
	res := GetProfileResponse{
		UserId:   id,
		Birthday: "2023-10-01",
		Address:  result,
		Phone:    "001",
		Sex:      "Male",
		Age:      20,
	}

	return &res, nil
}

func (profileServer) UpdateProfile(ctx context.Context, req *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	res := UpdateProfileResponse{
		Result: req.Address,
	}

	return &res, nil
}
