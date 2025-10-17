package user

import "boilerplate-echogo-dida/configs"

type UserRepository interface {
	GetUsersRepository() ([]Users, error)
	UpdateUserRepository(user Users) error
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (s *userRepo) GetUsersRepository() ([]Users, error) {
	var users []Users
	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userRepo) UpdateUserRepository(user Users) error {
	if err := configs.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
