package entities

import (
	"financial-api/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model
// swagger:model User
type User struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `gorm:"uniqueIndex" json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Password  string    `json:"password"`
	Accounts  []Account `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE" json:"accounts"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {

	// Generate password hash
	hash, err := utils.GenerateFromString(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	// Generate UUID
	u.UUID = uuid.New()

	return
}

func (u *User) Seed(db *gorm.DB) error {

	err := db.Create(&User{
		Name:     "John Doe",
		Email:    "jonh@doe.com",
		Password: "123456",
	}).Error

	if err != nil {
		return err
	}

	return nil
}
