package auth

import "demoecho/pkg/models"

type AuthRepository interface {
	Register(u models.User) (models.User, error)
	GetUserRepositoryByEmail(email string) (models.User, error)
}

type authRepository struct{}

func NewAuthRepo() AuthRepository {

	return &authRepository{}
}
