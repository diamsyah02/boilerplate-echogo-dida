package user

import "boilerplate-echogo-dida/configs"

func GetUsersRepository() ([]Users, error) {
	var users []Users
	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUserRepository(user Users) error {
	if err := configs.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
