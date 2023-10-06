package utils

import "golang.org/x/crypto/bcrypt"

type Password interface {
	HashPassword(password string) (string,error)
	ComparePassword(passwordHash, password string) (error)
}

func HashPassword(password string) (string,error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func ComparePassword(passwordHash, password string) (bool){
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
    return err == nil
}