package domain

import (
	"go-commerce-api/internal/payment/entity"
	"go-commerce-api/internal/product/domain"
	"time"

	"github.com/shopspring/decimal"
)

type Payment struct {
	ID          string
	PaymentCode string
	ProductID   string
	UserID      string
	Quantity    int
	TotalAmount decimal.Decimal
	Status      string
	PaymentURL  string
	Token       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Product     domain.Product
}

func PaymentDomainToEntity(paymentDomain Payment) entity.Payment {
	return entity.Payment{
		ID:          paymentDomain.ID,
		ProductID:   paymentDomain.ProductID,
		UserID:      paymentDomain.UserID,
		Quantity:    paymentDomain.Quantity,
		TotalAmount: paymentDomain.TotalAmount,
		Status:      paymentDomain.Status,
		PaymentURL:  paymentDomain.PaymentURL,
		Token:       paymentDomain.Token,
		CreatedAt:   paymentDomain.CreatedAt,
		UpdatedAt:   paymentDomain.UpdatedAt,
	}
}

func PaymentEntityToDomain(paymentEntity entity.Payment) Payment {
	return Payment{
		ID:          paymentEntity.ID,
		ProductID:   paymentEntity.ProductID,
		UserID:      paymentEntity.UserID,
		Quantity:    paymentEntity.Quantity,
		TotalAmount: paymentEntity.TotalAmount,
		Status:      paymentEntity.Status,
		PaymentURL:  paymentEntity.PaymentURL,
		Token:       paymentEntity.Token,
		CreatedAt:   paymentEntity.CreatedAt,
		UpdatedAt:   paymentEntity.UpdatedAt,
	}
}

func ListPaymentDomainToEntity(paymentDomains []Payment) []entity.Payment {
	listPaymentEntities := []entity.Payment{}
	for _, payment := range paymentDomains {
		listPaymentEntities = append(listPaymentEntities, PaymentDomainToEntity(payment))
	}
	return listPaymentEntities
}

func ListPaymentEntityToDomain(paymentEntities []entity.Payment) []Payment {
	listPaymentDomains := []Payment{}
	for _, payment := range paymentEntities {
		listPaymentDomains = append(listPaymentDomains, PaymentEntityToDomain(payment))
	}
	return listPaymentDomains
}
