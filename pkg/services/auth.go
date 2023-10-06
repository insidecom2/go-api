package services

import (
	"demoecho/pkg/models"
	"demoecho/pkg/repositories"
	"demoecho/pkg/utils"
	"errors"
)



type LoginBody struct {
	Email string ` validate:"email,required"`
	Password string ` validate:"required"`
}

type AuthService interface {
	RegisterAuth(u models.User) (models.User,error)
	LoginAuth(l LoginBody) (string,error)
}

type authService struct {}

var (
	authRepo  repositories.AuthRepository
)

func NewAuthService(repository repositories.AuthRepository)AuthService {
	authRepo = repository
	return &authService{}
}

func (s *authService) RegisterAuth(u models.User) (models.User,error){
	passwordHash ,errHash := utils.HashPassword(u.Password)
	if errHash != nil {
		return u,errHash
	}
	
	u.Password = passwordHash
	result ,err := authRepo.Register(u)

	if err != nil {
		return u,err
	}
	return result,nil

}

func  (s *authService) LoginAuth(l LoginBody) (string,error) {
	u,err := authRepo.GetUserRepositoryByEmail(l.Email)
	if err != nil {
		return "",err
	}

	isPasswordMatch := utils.ComparePassword(u.Password, l.Password)
	if !isPasswordMatch {
		return "",errors.New("Invalid Login")
	}

	token,e := utils.GenerateToken(u.Email)
	if e != nil {
		return "",errors.New("Invalid Login")
	}
	return token,nil
}