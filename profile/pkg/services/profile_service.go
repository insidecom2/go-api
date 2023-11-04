package services

import (
	context "context"
	"profile/pkg/repositories"
	"strconv"
)

var (
	profileRepo repositories.ProfileRepo
)

type profileServer struct{}

func NewProfileServer(repository repositories.ProfileRepo) ProfileServer {
	profileRepo = repository
	return profileServer{}
}

func (profileServer) mustEmbedUnimplementedProfileServer() {}

func (profileServer) GetProfile(ctx context.Context, in *GetProfileRequest) (*GetProfileResponse, error) {
	id, _ := strconv.ParseInt(in.UserId, 10, 64)
	profile, err := profileRepo.GetProfile(id)
	res := GetProfileResponse{}
	if err != nil {
		return &res, err
	}

	res.UserId = profile.UserId

	return &res, nil
}

func (profileServer) UpdateProfile(ctx context.Context, req *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	res := UpdateProfileResponse{
		Result: req.Address,
	}

	return &res, nil
}
