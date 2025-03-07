package entity

import (
	"go-commerce-api/internal/product/entity"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Payment struct {
	ID          string          `gorm:"type:varchar(36);primaryKey"`
	ProductID   string          `gorm:"type:varchar(36);not null;index"`
	UserID      string          `gorm:"type:varchar(36);not null"`
	PaymentCode string          `gorm:"type:varchar(36);not null;unique"`
	Quantity    int             `gorm:"not null"`
	TotalAmount decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	Status      string          `gorm:"type:enum('deny', 'success', 'cancel', 'expire', 'pending');default:'pending'"`
	PaymentURL  string          `gorm:"type:text"`
	Token       string          `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Product     entity.Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return nil
}
