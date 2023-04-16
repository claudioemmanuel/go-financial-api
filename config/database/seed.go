package database

import (
	"financial-api/domain/entities"
	"fmt"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {

	user := entities.User{
		Name:     "John Doe",
		Email:    "john@doe.com",
		Password: "123456",
	}

	err := db.Create(&user).Error
	if err != nil {
		fmt.Printf("Could not seed user entity \n")
		return err
	}

	account := entities.Account{
		OwnerID: user.ID,
		Balance: 1000,
	}

	err = db.Create(&account).Error
	if err != nil {
		fmt.Printf("Could not seed account entity \n")
		return err
	}

	fmt.Printf("Seeded entities successfully \n")

	return nil
}
