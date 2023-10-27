package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	UserId   int64     `json:"user_id"  size:"10" validate:"required" gorm:"unique;not null"`
	BirthDay time.Time `json:"birthday"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	Sex      string    `json:"sex"`
	Age      int16     `json:"age"`
}
