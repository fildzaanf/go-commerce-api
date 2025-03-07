package dto

import (
	"go-commerce-api/internal/product/domain"

	"github.com/shopspring/decimal"
)

type (
	CreateProductRequest struct {
		Name        string          `json:"name" form:"name"`
		Description string          `json:"description" form:"description"`
		Price       decimal.Decimal `json:"price" form:"price"`
		Stock       int             `json:"stock" form:"stock"`
		ImageURL    string          `json:"image_url" form:"image_url"`
	}

	UpdateProductRequest struct {
		Name        string          `json:"name" form:"name"`
		Description string          `json:"description" form:"description"`
		Price       decimal.Decimal `json:"price" form:"price"`
		Stock       int             `json:"stock" form:"stock"`
		ImageURL    string          `json:"image_url" form:"image_url"`
	}
)

func CreateProductRequestToDomain(request CreateProductRequest, userID string) domain.Product {
	return domain.Product{
		UserID:      userID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		ImageURL:    request.ImageURL,
	}
}

func UpdateProductRequestToDomain(request UpdateProductRequest, product *domain.Product) {
	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Description != "" {
		product.Description = request.Description
	}
	if request.Price.GreaterThan(decimal.NewFromInt(0)){
		product.Price = request.Price
	}
	if request.Stock >= 0 {
		product.Stock = request.Stock
	}
	if request.ImageURL != "" {
		product.ImageURL = request.ImageURL
	}
}
