package repositories

import (
	"demoecho/pkg/database"
	"demoecho/pkg/models"
)

type repository struct {}

type UserRepository interface {
	GetUserRepository(id string) (models.User, error)
}


func NewUserRepo() UserRepository {

	return &repository{}
}

func (r *repository) GetUserRepository(id string) (models.User, error) {

	var user = models.User{}

	result := database.DB.Where("ID = ?", id).First(&user)

	if result.Error != nil {
		return  user, result.Error
	}
	return user, nil
}
