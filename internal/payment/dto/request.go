package dto

import "go-commerce-api/internal/payment/domain"

type CreatePaymentRequest struct {
	ProductID string `json:"product_id" form:"product_id"`
	Quantity  int    `json:"quantity" form:"quantity"`
}

func CreatePaymentRequestToDomain(request CreatePaymentRequest) domain.Payment {
	return domain.Payment{
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
	}
}
