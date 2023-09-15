package services

import (
	"demoecho/pkg/interfaces"
)

type UserService interface {
	GetUserService(id string) string
	CreateUserService(u *interfaces.User) *interfaces.User
}

type services struct{}

func NewUserService() UserService {
	return &services{}
}


func (*services) GetUserService(id string)  string {

	return id

}

func (*services) CreateUserService(u *interfaces.User) *interfaces.User {

	return u

}
