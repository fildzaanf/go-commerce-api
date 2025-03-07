package repository

import (
	"errors"
	"go-commerce-api/internal/product/domain"

	"gorm.io/gorm"
)

type productCommandRepository struct {
	db *gorm.DB
}

func NewProductCommandRepository(db *gorm.DB) ProductCommandRepositoryInterface {
	return &productCommandRepository{
		db: db,
	}
}

func (pcr *productCommandRepository) CreateProduct(product domain.Product) (domain.Product, error) {
	tx := pcr.db.Begin()
	if tx.Error != nil {
		return domain.Product{}, tx.Error
	}

	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return domain.Product{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (pcr *productCommandRepository) UpdateProductByID(id string, product domain.Product) (domain.Product, error) {
	tx := pcr.db.Begin()
	if tx.Error != nil {
		return domain.Product{}, tx.Error
	}

	existingProduct := domain.Product{}
	if err := tx.First(&existingProduct, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return domain.Product{}, errors.New("product not found")
	}

	if err := tx.Model(&existingProduct).Updates(product).Error; err != nil {
		tx.Rollback()
		return domain.Product{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return domain.Product{}, err
	}

	return existingProduct, nil
}

func (pcr *productCommandRepository) DeleteProductByID(id string) error {
	tx := pcr.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("id = ?", id).Delete(&domain.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
