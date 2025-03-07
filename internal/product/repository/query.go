package repository

import (
	"errors"
	"go-commerce-api/internal/product/domain"
	"go-commerce-api/internal/product/entity"

	"gorm.io/gorm"
)

type productQueryRepository struct {
	db *gorm.DB
}

func NewProductQueryRepository(db *gorm.DB) ProductQueryRepositoryInterface {
	return &productQueryRepository{
		db: db,
	}
}

func (pr *productQueryRepository) GetProductByID(id string) (domain.Product, error) {
	var product entity.Product
	result := pr.db.Where("id = ?", id).First(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Product{}, errors.New("product not found")
		}
		return domain.Product{}, result.Error
	}

	return domain.ProductEntityToDomain(product), nil
}

func (pr *productQueryRepository) GetAllProducts() ([]domain.Product, error) {
	var products []entity.Product
	result := pr.db.Find(&products)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("no products found")
		}
		return nil, result.Error
	}

	return domain.ListProductEntityToDomain(products), nil
}
