package test

import (
	"testing"

	"financial-api/domain/entities"
	"financial-api/infrastructure/persistence/gorm"

	h "financial-api/tests"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db := h.SetupTestDB(t)
	defer h.CleanTestDB(t, db)

	userRepo := gorm.NewUserRepositoryGorm(db)

	user := entities.User{
		Email:    "test@example.com",
		Password: "password",
	}

	err := userRepo.Create(&user)
	assert.Nil(t, err)
	assert.NotZero(t, user.ID)
}

func TestUserRepository_FindAll(t *testing.T) {
	db := h.SetupTestDB(t)
	defer h.CleanTestDB(t, db)

	userRepo := gorm.NewUserRepositoryGorm(db)

	user := entities.User{
		Email:    "test@example.com",
		Password: "password",
	}

	err := userRepo.Create(&user)
	assert.Nil(t, err)
	assert.NotZero(t, user.ID)
}
