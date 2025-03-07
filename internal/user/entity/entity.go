package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null"`
	Password  string `gorm:"type:text;not null"`
	Role      string `gorm:"type:enum('user','seller','buyer');default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()

	if u.Role == "" {
		u.Role = "user"
	}

	return nil
}
