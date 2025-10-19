package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(req Users) (string, error)
	Register(req Users) error
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Login(req Users) (string, error) {
	result, found, err := s.repo.Login(req)
	if err != nil {
		return "", err
	}
	if !found {
		return "", nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password)); err != nil {
		return "", nil
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": result.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s *authService) Register(req Users) error {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(password)
	err = s.repo.Register(req)
	if err != nil {
		return err
	}
	return nil
}
