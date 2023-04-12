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
	UUID      uuid.UUID `gorm:"unique" json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

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
