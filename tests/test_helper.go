package test

import (
	"financial-api/domain/entities"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("Failed to connect to the test database: %v", err.Error())
	}

	entities := []interface{}{
		&entities.User{},
		&entities.Account{},
	}

	for _, entity := range entities {
		err := db.AutoMigrate(entity)
		if err != nil {
			t.Fatalf("Failed to migrate the test database: %v", err.Error())
		}
	}

	return db
}

func CleanTestDB(t *testing.T, db *gorm.DB) {
	err := db.Migrator().DropTable(&entities.User{})
	if err != nil {
		t.Fatalf("Failed to drop the test database: %v", err.Error())
	}
}
