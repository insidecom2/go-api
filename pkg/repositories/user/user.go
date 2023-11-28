package user

import (
	"demoecho/pkg/database"
	"demoecho/pkg/models"
)

func (r *repository) GetUserRepository(id string) (models.User, error) {

	var user = models.User{}

	result := database.DB.Where("ID = ?", id).First(&user)

	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
