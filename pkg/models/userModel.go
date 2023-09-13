package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" size:"255"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"passoword"`
	Token    string `json:"token"`
}
