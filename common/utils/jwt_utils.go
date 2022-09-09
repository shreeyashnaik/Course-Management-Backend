package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/shreeyashnaik/Course-Management-Backend/config"
)

// Creating JWT Token
func CreateJWTToken(id, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 7 * 24).Unix()

	t, err := token.SignedString([]byte(config.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return t, nil
}
