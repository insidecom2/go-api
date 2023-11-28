package services

import (
	"demoecho/pkg/models"
	repositories "demoecho/pkg/repositories/user"
)

type UserService interface {
	GetUserService(id string) (models.User, error)
}

type userServices struct{}

var (
	userRepo repositories.UserRepository
)

func NewUserService(repositories repositories.UserRepository) UserService {
	userRepo = repositories
	return &userServices{}
}

func (s *userServices) GetUserService(id string) (models.User, error) {

	result, err := userRepo.GetUserRepository(id)

	return result, err
}
