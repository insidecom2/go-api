package repositories

import (
	"profile/pkg/database"
	"profile/pkg/models"
)

type profileRepo struct{}

type ProfileRepo interface {
	GetProfile(id int64) (models.Profile, error)
	UpdateProfile(profile models.Profile) error
}

func NewProfileRepo() ProfileRepo {

	return &profileRepo{}
}

func (p *profileRepo) GetProfile(id int64) (models.Profile, error) {
	var profile = models.Profile{}

	result := database.DB.Where("user_id = ?", id).First(&profile)

	if result.Error != nil {
		return profile, result.Error
	}
	return profile, nil
}

func (p *profileRepo) UpdateProfile(profile models.Profile) error {
	return nil
}
