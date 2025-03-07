package repository

import (
	"errors"
	"go-commerce-api/internal/payment/domain"

	"gorm.io/gorm"
)

type paymentQueryRepository struct {
	db *gorm.DB
}

func NewPaymentQueryRepository(db *gorm.DB) PaymentQueryRepositoryInterface {
	return &paymentQueryRepository{
		db: db,
	}
}

func (pqr *paymentQueryRepository) GetPaymentByID(id string) (domain.Payment, error) {
	var payment domain.Payment
	result := pqr.db.Preload("Product").Where("id = ?", id).First(&payment)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Payment{}, errors.New("payment not found")
		}
		return domain.Payment{}, result.Error
	}

	return payment, nil
}


func (pqr *paymentQueryRepository) GetAllPayments(userID string) ([]domain.Payment, error) {
	var payments []domain.Payment
	result := pqr.db.Preload("Product").Where("user_id = ?", userID).Find(&payments)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("no payments found")
		}
		return nil, result.Error
	}

	return payments, nil
}
