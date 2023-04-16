package database

import (
	"financial-api/domain/entities"
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	// get all entities
	entities := []interface{}{
		&entities.User{},
		&entities.Account{},
	}

	// migrate all entities
	for _, entity := range entities {
		err := db.AutoMigrate(entity)
		if err != nil {
			fmt.Printf("Could not migrate %s entity \n", entity)
			return err
		}

		fmt.Printf("Migrated %s entity \n", entity)
	}

	return nil
}
