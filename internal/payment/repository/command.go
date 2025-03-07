package repository

import (
	"go-commerce-api/internal/payment/domain"

	"gorm.io/gorm"
)

type paymentCommandRepository struct {
	db *gorm.DB
}

func NewPaymentCommandRepository(db *gorm.DB) PaymentCommandRepositoryInterface {
	return &paymentCommandRepository{
		db: db,
	}
}

func (pcr *paymentCommandRepository) CreatePayment(payment domain.Payment) (domain.Payment, error) {
	tx := pcr.db.Begin()
	if tx.Error != nil {
		return domain.Payment{}, tx.Error
	}

	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		return domain.Payment{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return domain.Payment{}, err
	}

	return payment, nil
}


func (pcr *paymentCommandRepository) UpdatePaymentStatusByID(id string, status string) error {
	tx := pcr.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Model(&domain.Payment{}).
		Where("id = ?", id).
		Update("status", status).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
