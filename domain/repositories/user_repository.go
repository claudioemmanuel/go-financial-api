package repositories

import "financial-api/domain/entities"

type UserRepository interface {
	FindAll() ([]entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(user *entities.User) error
}
