package services

import (
	"demoecho/pkg/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUserService(t *testing.T) {
	testService := NewUserService()

	err := testService.ValidateUserService(nil)

	assert.Equal(t, err.Error(), "Data not found!!!")
}

func TestValidateUserServiceWithName(t *testing.T) {
	testService := NewUserService()
	u := &interfaces.User{
		Name:    "",
		Surname: "",
	}
	err := testService.ValidateUserService(u)

	assert.Equal(t, err.Error(), "Name not found!!!")
}

func TestValidateUserServiceWithSurname(t *testing.T) {
	testService := NewUserService()
	u := &interfaces.User{
		Name:    "Thom",
		Surname: "",
	}
	err := testService.ValidateUserService(u)

	assert.Equal(t, err.Error(), "Surname not found!!!")
}

func TestValidateUserServicePass(t *testing.T) {
	testService := NewUserService()
	u := &interfaces.User{
		Name:    "Thom",
		Surname: "KKKK",
	}
	err := testService.ValidateUserService(u)

	assert.Equal(t, err, nil)
}

func TestGetUserService(t *testing.T) {
	testService := NewUserService()
	u := &interfaces.User{
		Name:    "Thom",
		Surname: "KKKK",
	}

	value := testService.GetUserService(u)
	assert.Equal(t, value, u)
}

func TestGetUserServiceOnlyName(t *testing.T) {
	testService := NewUserService()
	u := &interfaces.User{
		Name:    "Thom",
		Surname: "",
	}

	value := testService.GetUserService(u)
	assert.Equal(t, value, u)
}
