package user

type UserService interface {
	GetUsers() ([]Users, error)
	UpdateUser(user Users) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUsers() ([]Users, error) {
	result, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) UpdateUser(user Users) error {
	err := s.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
