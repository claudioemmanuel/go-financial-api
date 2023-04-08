package services

import (
	"financial-api/domain/entities"
	"financial-api/domain/repositories"
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

func (s *UserService) Delete(user *entities.User) error {
	return s.userRepository.Delete(user)
}
