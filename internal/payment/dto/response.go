package dto

import (
	"go-commerce-api/internal/payment/domain"
	product "go-commerce-api/internal/product/dto"
	"time"

	"github.com/shopspring/decimal"
)

type PaymentResponse struct {
	ID          string                  `json:"id"`
	ProductID   string                  `json:"product_id"`
	UserID      string                  `json:"user_id"`
	PaymentCode string                  `json:"payment_code"`
	Quantity    int                     `json:"quantity"`
	TotalAmount decimal.Decimal         `json:"total_amount"`
	Status      string                  `json:"status"`
	PaymentURL  string                  `json:"payment_url"`
	Token       string                  `json:"token"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
	Product     product.ProductResponse `json:"product"`
}

func PaymentDomainToResponse(payment domain.Payment) PaymentResponse {
	return PaymentResponse{
		ID:          payment.ID,
		ProductID:   payment.ProductID,
		UserID:      payment.UserID,
		PaymentCode: payment.PaymentCode,
		Quantity:    payment.Quantity,
		TotalAmount: payment.TotalAmount,
		Status:      payment.Status,
		PaymentURL:  payment.PaymentURL,
		Token:       payment.Token,
		CreatedAt:   payment.CreatedAt,
		UpdatedAt:   payment.UpdatedAt,
		Product:     product.ProductDomainToResponse(payment.Product),
	}
}

func ListPaymentDomainToResponse(payments []domain.Payment) []PaymentResponse {
	paymentResponses := make([]PaymentResponse, len(payments))
	for i, payment := range payments {
		paymentResponses[i] = PaymentDomainToResponse(payment)
	}
	return paymentResponses
}
