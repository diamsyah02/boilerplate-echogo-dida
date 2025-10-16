package user

func GetUsersService() ([]Users, error) {
	result, err := GetUsersRepository()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUserService(user Users) error {
	err := UpdateUserRepository(user)
	if err != nil {
		return err
	}
	return nil
}
