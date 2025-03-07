package repository

import "go-commerce-api/internal/product/domain"

type ProductCommandRepositoryInterface interface {
	CreateProduct(product domain.Product) (domain.Product, error)
	UpdateProductByID(id string, product domain.Product) (domain.Product, error)
	DeleteProductByID(id string) error
}

type ProductQueryRepositoryInterface interface {
	GetProductByID(id string) (domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
}
