package services

import (
	"context"
)

type ProfileService interface {
	GetProfile(user_id string) (*GetProfileResponse, error)
}

type profileService struct {
	profileClient ProfileClient
}

func NewProfileService(profileClient ProfileClient) ProfileService {
	return profileService{profileClient}
}

func (p profileService) GetProfile(user_id string) (*GetProfileResponse, error) {

	req := GetProfileRequest{
		UserId: user_id,
	}

	res, err := p.profileClient.GetProfile(context.Background(), &req)
	if err != nil {
		return res, err
	}

	return res, nil
}
