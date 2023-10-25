package middlewares

import (
	"demoecho/pkg/response"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		secretKey := []byte(os.Getenv("JWT_KEY"))
		authHeader := strings.Split(c.Request().Header.Get("Authorization"), " ")
		if len(authHeader) < 2 {
			return c.JSON(http.StatusUnauthorized, response.ResponseUnAuth("Invalid token"))
		}
		tokenString := authHeader[1]

		token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.ResponseUnAuth("Invalid token"))
		}

		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			c.Set("email", claims.Email)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, response.ResponseUnAuth("Invalid token"))
	}
}

func RefreshToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		secretKey := []byte(os.Getenv("JWT_REFRESH_KEY"))
		authHeader := strings.Split(c.Request().Header.Get("Authorization"), " ")
		if len(authHeader) < 2 {
			return c.JSON(http.StatusUnauthorized, response.ResponseUnAuth("Invalid token"))
		}
		tokenString := authHeader[1]

		token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.ResponseUnAuth("Invalid token"))
		}
		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			c.Set("email", claims.Email)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, response.ResponseUnAuth("Invalid token"))

	}

}
