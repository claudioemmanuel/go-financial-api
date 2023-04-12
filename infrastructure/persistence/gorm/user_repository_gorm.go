package gorm

import (
	"financial-api/domain/entities"

	"gorm.io/gorm"
)

type userRepositoryGorm struct {
	db *gorm.DB
}

func NewUserRepositoryGorm(db *gorm.DB) *userRepositoryGorm {
	return &userRepositoryGorm{
		db: db,
	}
}

func (r *userRepositoryGorm) FindAll() ([]entities.User, error) {
	var users []entities.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepositoryGorm) FindByID(id int) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *userRepositoryGorm) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryGorm) Update(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *userRepositoryGorm) Delete(user *entities.User) error {
	return r.db.Delete(user).Error
}

func (r *userRepositoryGorm) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
