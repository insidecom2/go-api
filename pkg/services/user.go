package services

import (
	"demoecho/pkg/models"
	"demoecho/pkg/repositories"
	"demoecho/pkg/utils"
)

type UserService interface {
	GetUserService(id string) (models.User,error)
	CreateUserService(u models.User) (models.User,error)
}

type services struct{}

var (
	userRepo repositories.UserRepository
)

func NewUserService(repositories repositories.UserRepository) UserService {
	userRepo = repositories
	return &services{}
}


func (s *services) GetUserService(id string)  (models.User,error) {

	result,err := userRepo.GetUserRepository(id)
	
	return result,err
}

func (s *services) CreateUserService(u models.User) (models.User,error){
	passwordHash ,errHash := utils.HashPassword(u.Password)
	if errHash != nil {
		return u,errHash
	}
	
	u.Password = passwordHash
	result ,err := userRepo.CreateUserRepository(u)

	if err != nil {
		return u,err
	}
	return result,nil

}
