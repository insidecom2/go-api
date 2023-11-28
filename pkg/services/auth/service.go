package auth

import (
	"demoecho/pkg/models"
	repositories "demoecho/pkg/repositories/auth"
	"demoecho/pkg/requests"
)

type AuthService interface {
	RegisterAuth(u models.User) (models.User, error)
	LoginAuth(l requests.LoginBody) (string, string, error)
	RefreshAuth(email string) (string, string, error)
}

func NewAuthService(repository repositories.AuthRepository) AuthService {
	authRepo = repository
	return &authService{}
}
