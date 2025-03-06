package entity

import (
	"go-commerce-api/internal/product/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID          string `gorm:"type:varchar(36);primaryKey"`
	ProductID   string `gorm:"type:varchar(36);not null;index"`
	Quantity    int    `gorm:"not null"`
	TotalAmount int    `gorm:"not null"`
	Status      string `gorm:"type:enum('deny', 'success', 'cancel', 'expire', 'pending');default:'pending'"`
	PaymentURL  string `gorm:"type:text"`
	Token       string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Product     entity.Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return nil
}
