package user

import "demoecho/pkg/models"

type repository struct{}

type UserRepository interface {
	GetUserRepository(id string) (models.User, error)
}

func NewUserRepo() UserRepository {

	return &repository{}
}
