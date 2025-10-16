package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(req Users) (string, error) {
	result, err := LoginRepository(req)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password)); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": result.Username,
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func RegisterService(req Users) error {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(password)
	err = RegisterRepository(req)
	if err != nil {
		return err
	}
	return nil
}
