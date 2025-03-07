package service

import "go-commerce-api/internal/payment/domain"

type PaymentCommandServiceInterface interface {
	CreatePayment(payment domain.Payment, userID string) (domain.Payment, error)
	UpdatePaymentStatusByID(id, status string) error
}

type PaymentQueryServiceInterface interface {
	GetPaymentByID(id string) (domain.Payment, error)
	GetAllPayments(userID string) ([]domain.Payment, error)
}