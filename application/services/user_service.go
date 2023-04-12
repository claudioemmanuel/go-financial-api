package services

import (
	"financial-api/domain/entities"
	"financial-api/domain/repositories"
	"financial-api/utils"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) FindAll() ([]entities.User, error) {
	return s.userRepository.FindAll()
}

func (s *UserService) Create(user *entities.User) error {
	return s.userRepository.Create(user)
}

func (s *UserService) Update(user *entities.User) error {
	return s.userRepository.Update(user)
}

func (s *UserService) Delete(id int) error {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return err
	}
	return s.userRepository.Delete(user)
}

func (s *UserService) Login(email, password string) (*entities.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := utils.CompareHashAndString(password, []byte(user.Password)); err != nil {
		return nil, err
	}

	return user, nil
}
