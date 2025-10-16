package auth

import "boilerplate-echogo-dida/configs"

func LoginRepository(req Users) (Users, error) {
	if err := configs.DB.Where("username = ?", req.Username).First(&req).Error; err != nil {
		return Users{}, err
	}
	return req, nil
}

func RegisterRepository(req Users) error {
	if err := configs.DB.Create(&req).Error; err != nil {
		return err
	}
	return nil
}
