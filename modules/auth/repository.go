package auth

import (
	"boilerplate-echogo-dida/configs"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(req Users) (Users, bool, error)
	Register(req Users) error
}

type authRepository struct{}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (s *authRepository) Login(req Users) (Users, bool, error) {
	if err := configs.DB.Where("username = ?", req.Username).First(&req).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Users{}, false, nil
		}
		return Users{}, false, err
	}
	return req, true, nil
}

func (s *authRepository) Register(req Users) error {
	if err := configs.DB.Create(&req).Error; err != nil {
		return err
	}
	return nil
}
