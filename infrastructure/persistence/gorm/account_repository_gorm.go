package gorm

import (
	"financial-api/domain/entities"

	"gorm.io/gorm"
)

type accountRepositoryGorm struct {
	db *gorm.DB
}

func NewAccountRepositoryGorm(db *gorm.DB) *accountRepositoryGorm {
	return &accountRepositoryGorm{
		db: db,
	}
}

func (r *accountRepositoryGorm) FindAll() ([]entities.Account, error) {
	var accounts []entities.Account
	err := r.db.Find(&accounts).Error
	return accounts, err
}

func (r *accountRepositoryGorm) FindByID(id int) (*entities.Account, error) {
	var account entities.Account
	err := r.db.Where("id = ?", id).First(&account).Error
	return &account, err
}

func (r *accountRepositoryGorm) Create(account *entities.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepositoryGorm) Update(account *entities.Account) error {
	return r.db.Save(account).Error
}

func (r *accountRepositoryGorm) Delete(account *entities.Account) error {
	return r.db.Delete(account).Error
}

func (r *accountRepositoryGorm) FindByOwnerID(ownerID int) (*entities.Account, error) {
	var account entities.Account
	err := r.db.Where("owner_id = ?", ownerID).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}
