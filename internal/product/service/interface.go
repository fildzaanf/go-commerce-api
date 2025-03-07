package service

import (
	"go-commerce-api/internal/product/domain"
	"mime/multipart"
)

type ProductCommandServiceInterface interface {
	CreateProduct(product domain.Product, image *multipart.FileHeader) (domain.Product, error)
	UpdateProductByID(id string, product domain.Product, image *multipart.FileHeader) (domain.Product, error)
	DeleteProductByID(id string) error
}

type ProductQueryServiceInterface interface {
	GetProductByID(id string) (domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
}
