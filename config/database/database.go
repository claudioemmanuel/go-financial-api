package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type DatabaseConfig struct{}

func Connect() (*gorm.DB, error) {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo", HOST, USER, PASS, DBNAME)

	var maxAttempts = 10

	// Try to connect to the database for 10 times
	for i := 0; i < maxAttempts; i++ {
		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			db = connection
			break
		}
		log.Printf("Could not connect to the database, attempt: %d", i+1)
		time.Sleep(5 * time.Second)
	}

	if db == nil {
		return nil, fmt.Errorf("failed to connect to the database after %d attempts", maxAttempts)
	}

	// Ping the database
	psqlDB, _ := db.DB()
	err := psqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping database: %v", err)
	}

	// Migrate the database
	err = Migrate(db)
	if err != nil {
		return nil, fmt.Errorf("could not migrate database: %v", err)
	}

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}

func Close(db *gorm.DB) error {
	psqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Could not get database: ", err)
	}

	err = psqlDB.Close()
	if err != nil {
		log.Fatal("Could not close database: ", err)
	}

	return nil
}
