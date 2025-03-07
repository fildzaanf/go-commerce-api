package repository

import "go-commerce-api/internal/payment/domain"

type PaymentCommandRepositoryInterface interface {
	CreatePayment(payment domain.Payment) (domain.Payment, error)
	UpdatePaymentStatusByID(id, status string) error
}

type PaymentQueryRepositoryInterface interface {
	GetAllPayments(userID string) ([]domain.Payment, error)
	GetPaymentByID(id string) (domain.Payment, error)
}
