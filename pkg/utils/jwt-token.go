package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type GenerateTokenProp struct {
	Email string `json:"email"`
	Id    uint   `json:"id"`
}

func GenerateToken(prop GenerateTokenProp) (string, string, error) {

	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	var jwtRefreshKey = []byte(os.Getenv("JWT_REFRESH_KEY"))
	var err error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = prop.Email
	atClaims["id"] = prop.Id
	atClaims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, errToken := at.SignedString([]byte(jwtKey))
	refreshToken, errRefresh := at.SignedString([]byte(jwtRefreshKey))

	if errToken != nil || errRefresh != nil {
		return "", "", err
	}
	return token, refreshToken, nil

}

// Get retrieves data from the context.
func GetContext(ctx echo.Context, key string) string {
	return fmt.Sprint(ctx.Get(key))
}
