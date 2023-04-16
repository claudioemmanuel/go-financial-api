package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Account model
// swagger:model Account
type Account struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `gorm:"uniqueIndex" json:"uuid"`
	OwnerID   uint      `json:"owner_id"`
	User      User      `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE" json:"user"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (a *Account) BeforeCreate(db *gorm.DB) (err error) {

	// Generate UUID
	a.UUID = uuid.New()

	return
}
