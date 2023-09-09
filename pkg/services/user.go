package services

import (
	"demoecho/pkg/interfaces"
	"errors"
)

type UserService interface {
	GetUserService(u *interfaces.User) *interfaces.User
	ValidateUserService(u *interfaces.User) error
}

type services struct{}

func NewUserService() UserService {
	return &services{}
}

func (*services) ValidateUserService(u *interfaces.User) error {
	if u == nil {
		err := errors.New("Data not found!!!")
		return err
	} else if u.Name == "" {
		err := errors.New("Name not found!!!")
		return err
	} else if u.Surname == "" {
		err := errors.New("Surname not found!!!")
		return err
	}

	return nil
}

func (*services) GetUserService(u *interfaces.User) *interfaces.User {

	return u

}
