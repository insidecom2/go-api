package repositories

import (
	"demoecho/pkg/database"
	"demoecho/pkg/models"
)

type repository struct {}

type UserRepository interface {
	GetUserRepository(id string) (models.User, error) 
	CreateUserRepository(u models.User) (models.User, error)
}


func NewUserRepo() UserRepository {

	return &repository{}
}

func (r *repository) CreateUserRepository(u models.User) (models.User, error) {

	result := database.DB.Create(&u)

	if result.Error != nil {
		return u ,result.Error
	}

	return u,nil
}

func (r *repository) GetUserRepository(id string) (models.User, error) {

	var user = models.User{}

	result := database.DB.Where("ID = ?", id).First(&user)

	if result.Error != nil {
		return  user, result.Error
	}
	return user, nil
}