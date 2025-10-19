package user

import "boilerplate-echogo-dida/pkg/configs"

type UserRepository interface {
	GetUsers() ([]Users, error)
	UpdateUser(user Users) error
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (s *userRepo) GetUsers() ([]Users, error) {
	var users []Users
	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userRepo) UpdateUser(user Users) error {
	if err := configs.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
