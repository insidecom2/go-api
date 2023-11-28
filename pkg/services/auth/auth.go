package auth

import (
	"demoecho/pkg/constants"
	"demoecho/pkg/models"
	repositories "demoecho/pkg/repositories/auth"
	"demoecho/pkg/requests"
	"demoecho/pkg/utils"
	"errors"
)

type authService struct{}

var (
	authRepo repositories.AuthRepository
)

func (s *authService) RegisterAuth(u models.User) (models.User, error) {
	passwordHash, errHash := utils.HashPassword(u.Password)
	if errHash != nil {
		return u, errHash
	}

	u.Password = passwordHash
	result, err := authRepo.Register(u)

	if err != nil {
		return u, err
	}
	return result, nil

}

func (s *authService) LoginAuth(l requests.LoginBody) (string, string, error) {
	u, err := authRepo.GetUserRepositoryByEmail(l.Email)
	if err != nil {
		return "", "", err
	}

	isPasswordMatch := utils.ComparePassword(u.Password, l.Password)
	if !isPasswordMatch {
		return "", "", errors.New(constants.ErrMsg["UNAUTHORIZE"])
	}

	tokenProp := utils.GenerateTokenProp{
		Email: u.Email,
		Id:    u.ID,
	}

	token, refreshToken, e := utils.GenerateToken(tokenProp)
	if e != nil {
		return "", "", errors.New(constants.ErrMsg["UNAUTHORIZE"])
	}
	return token, refreshToken, nil
}

func (s *authService) RefreshAuth(email string) (string, string, error) {
	u, err := authRepo.GetUserRepositoryByEmail(email)
	if err != nil {
		return "", "", err
	}
	tokenProp := utils.GenerateTokenProp{
		Email: u.Email,
		Id:    u.ID,
	}
	token, refreshToken, e := utils.GenerateToken(tokenProp)
	if e != nil {
		return "", "", errors.New(constants.ErrMsg["UNAUTHORIZE"])
	}
	return token, refreshToken, nil
}
