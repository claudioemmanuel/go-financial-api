package test

import (
	"testing"

	"financial-api/domain/entities"
	"financial-api/infrastructure/persistence/gorm"

	h "financial-api/tests"

	"github.com/stretchr/testify/assert"
)

func TestAccountRepository_Create(t *testing.T) {
	db := h.SetupTestDB(t)
	defer h.CleanTestDB(t, db)

	accountRepo := gorm.NewAccountRepositoryGorm(db)

	user := entities.User{
		Email:    "test@example.com",
		Password: "password",
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
	assert.NotZero(t, user.ID)

	account := entities.Account{
		OwnerID: user.ID,
		Balance: 100,
	}

	err = accountRepo.Create(&account)
	assert.Nil(t, err)
	assert.NotZero(t, account.ID)
	assert.Equal(t, account.OwnerID, user.ID)
	assert.Equal(t, account.Balance, 100)
}

func TestAccountRepository_FindAll(t *testing.T) {
	db := h.SetupTestDB(t)
	defer h.CleanTestDB(t, db)

	accountRepo := gorm.NewAccountRepositoryGorm(db)

	user := entities.User{
		Email:    "test@example.com",
		Password: "password",
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
	assert.NotZero(t, user.ID)

	account := entities.Account{
		OwnerID: user.ID,
		Balance: 100,
	}

	err = accountRepo.Create(&account)
	assert.Nil(t, err)
	assert.NotZero(t, account.ID)

	accounts, err := accountRepo.FindAll()
	assert.Nil(t, err)
	assert.NotZero(t, len(accounts))
}
