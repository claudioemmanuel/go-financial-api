package database

import (
	"financial-api/domain/entities"
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	err := db.AutoMigrate(&entities.User{})
	if err != nil {
		fmt.Println("Could not migrate User entity")
		return err
	}

	return nil
}
