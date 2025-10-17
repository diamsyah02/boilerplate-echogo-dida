package user

type UserService interface {
	GetUsersService() ([]Users, error)
	UpdateUserService(user Users) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUsersService() ([]Users, error) {
	result, err := s.repo.GetUsersRepository()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) UpdateUserService(user Users) error {
	err := s.repo.UpdateUserRepository(user)
	if err != nil {
		return err
	}
	return nil
}
