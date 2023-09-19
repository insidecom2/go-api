package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name    string `json:"name"  size:"255" validate:"required"`
	Surname string `json:"surname" validate:"required"`
	Email string `json:"email" validate:"email,required" gorm:"unique,{Over here}not null"`
	Password string `json:"password" validate:"required"`
	Token string `json:"token"`
}
