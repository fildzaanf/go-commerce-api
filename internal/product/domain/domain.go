package domain

import (
	"go-commerce-api/internal/product/entity"
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID          string
	UserID      string
	Name        string
	Description string
	Price       decimal.Decimal
	Stock       int
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// mapper
func ProductDomainToEntity(productDomain Product) entity.Product {
	return entity.Product{
		ID:          productDomain.ID,
		UserID:      productDomain.UserID,
		Name:        productDomain.Name,
		Description: productDomain.Description,
		Price:       productDomain.Price,
		Stock:       productDomain.Stock,
		ImageURL:    productDomain.ImageURL,
		CreatedAt:   productDomain.CreatedAt,
		UpdatedAt:   productDomain.UpdatedAt,
		DeletedAt:   productDomain.DeletedAt,
	}
}

func ProductEntityToDomain(productEntity entity.Product) Product {
	return Product{
		ID:          productEntity.ID,
		UserID:      productEntity.UserID,
		Name:        productEntity.Name,
		Description: productEntity.Description,
		Price:       productEntity.Price,
		Stock:       productEntity.Stock,
		ImageURL:    productEntity.ImageURL,
		CreatedAt:   productEntity.CreatedAt,
		UpdatedAt:   productEntity.UpdatedAt,
		DeletedAt:   productEntity.DeletedAt,
	}
}

func ListProductDomainToEntity(productDomains []Product) []entity.Product {
	listProductEntities := []entity.Product{}
	for _, product := range productDomains {
		listProductEntities = append(listProductEntities, ProductDomainToEntity(product))
	}
	return listProductEntities
}

func ListProductEntityToDomain(productEntities []entity.Product) []Product {
	listProductDomains := []Product{}
	for _, product := range productEntities {
		listProductDomains = append(listProductDomains, ProductEntityToDomain(product))
	}
	return listProductDomains
}
