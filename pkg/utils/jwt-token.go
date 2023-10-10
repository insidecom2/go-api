package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


func GenerateToken(email string) (string,string,error) {
	
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	var jwtRefreshKey = []byte(os.Getenv("JWT_REFRESH_KEY"))
	var err error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, errToken := at.SignedString([]byte(jwtKey))
	refreshToken, errRefresh := at.SignedString([]byte(jwtRefreshKey))

	if errToken != nil || errRefresh != nil {
		return "", "",err
	}
	return token,refreshToken, nil

}