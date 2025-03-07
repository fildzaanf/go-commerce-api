package dto

import (
	"go-commerce-api/internal/product/domain"
	"time"

	"github.com/shopspring/decimal"
)

type (
	ProductResponse struct {
		ID          string          `json:"id"`
		UserID      string          `json:"user_id"`
		Name        string          `json:"name"`
		Description string          `json:"description"`
		Price       decimal.Decimal `json:"price"`
		Stock       int             `json:"stock"`
		ImageURL    string          `json:"image_url"`
		CreatedAt   time.Time       `json:"created_at"`
		UpdatedAt   time.Time       `json:"updated_at"`
	}
)

func ProductDomainToResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		UserID:      product.UserID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		ImageURL:    product.ImageURL,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ListProductDomainToResponse(products []domain.Product) []ProductResponse {
	productResponses := make([]ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = ProductDomainToResponse(product)
	}
	return productResponses
}
