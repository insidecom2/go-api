package repositories

import (
	"demoecho/pkg/database"
	"demoecho/pkg/models"
)

type AuthRepository interface {
	Register(u models.User) (models.User, error)
	GetUserRepositoryByEmail(email string) (models.User, error) 
}

type authRepository struct {}

func NewAuthRepo() AuthRepository {

	return &authRepository{}
}

func (r *authRepository) Register(u models.User) (models.User, error) {

	result := database.DB.Create(&u)

	if result.Error != nil {
		return u ,result.Error
	}

	return u,nil
}


func (r *authRepository) GetUserRepositoryByEmail(email string) (models.User, error) {

	var user = models.User{}

	result := database.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return  user, result.Error
	}
	return user, nil
}
