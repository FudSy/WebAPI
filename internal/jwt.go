package internal

import (
	"os"
	"time"

	"github.com/FudSy/WebApi/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userdata models.User) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userdata.UserName,                // Subject (user identifier)
		"aud": userdata.Role,                    // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})

	tokenString, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
