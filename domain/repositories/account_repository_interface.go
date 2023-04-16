package repositories

import "financial-api/domain/entities"

type AccountRepository interface {
	FindAll() ([]entities.Account, error)
	FindByID(id int) (*entities.Account, error)
	Create(account *entities.Account) error
	Update(account *entities.Account) error
	Delete(account *entities.Account) error
	FindByOwnerID(ownerID int) (*entities.Account, error)
}
