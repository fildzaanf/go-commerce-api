package service

import (
	"errors"
	"go-commerce-api/internal/product/domain"
	"go-commerce-api/internal/product/repository"
	"go-commerce-api/pkg/cloud"
	"go-commerce-api/pkg/constant"
	"go-commerce-api/pkg/validator"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type productCommandService struct {
	productCommandRepository repository.ProductCommandRepositoryInterface
	productQueryRepository   repository.ProductQueryRepositoryInterface
}

func NewProductCommandService(pcr repository.ProductCommandRepositoryInterface, pqr repository.ProductQueryRepositoryInterface) ProductCommandServiceInterface {
	return &productCommandService{
		productCommandRepository: pcr,
		productQueryRepository:   pqr,
	}
}

func (pcs *productCommandService) CreateProduct(product domain.Product, image *multipart.FileHeader) (domain.Product, error) {

	if image != nil {
		imageURL, errUpload := cloud.UploadImageToS3(image)
		if errUpload != nil {
			return domain.Product{}, errUpload
		}
		product.ImageURL = imageURL
	}

	if err := validator.IsDataEmpty([]string{"name", "description", "image_url", "price", "stock"}, product.Name, product.Description, product.ImageURL, product.Price, product.Stock); err != nil {
		return domain.Product{}, err
	}

	if product.Price.LessThanOrEqual(decimal.NewFromInt(0)) {
		return domain.Product{}, errors.New(constant.ERROR_INVALID_PRICE)
	}

	if product.Stock < 0 {
		return domain.Product{}, errors.New(constant.ERROR_INVALID_STOCK)
	}

	if product.ID == "" {
		product.ID = uuid.New().String()
	}

	createdProduct, err := pcs.productCommandRepository.CreateProduct(product)
	if err != nil {
		return domain.Product{}, err
	}

	return createdProduct, nil
}
func (pcs *productCommandService) UpdateProductByID(id string, product domain.Product, image *multipart.FileHeader) (domain.Product, error) {
	existingProduct, err := pcs.productQueryRepository.GetProductByID(id)
	if err != nil {
		return domain.Product{}, errors.New(constant.ERROR_PRODUCT_NOT_FOUND)
	}

	if product.Name != "" {
		existingProduct.Name = product.Name
	}
	if product.Description != "" {
		existingProduct.Description = product.Description
	}
	if product.Price.GreaterThan(decimal.NewFromInt(0)) {
		existingProduct.Price = product.Price
	}
	if product.Stock >= 0 {
		existingProduct.Stock = product.Stock
	}

	if image != nil {
		imageURL, errUpload := cloud.UploadImageToS3(image)
		if errUpload != nil {
			return domain.Product{}, errUpload
		}
		existingProduct.ImageURL = imageURL
	}

	existingProduct.ID = id

	updatedProduct, err := pcs.productCommandRepository.UpdateProductByID(id, existingProduct)
	if err != nil {
		return domain.Product{}, err
	}

	return updatedProduct, nil
}


func (pcs *productCommandService) DeleteProductByID(id string) error {
	_, err := pcs.productQueryRepository.GetProductByID(id)
	if err != nil {
		return errors.New(constant.ERROR_PRODUCT_NOT_FOUND)
	}

	if err := pcs.productCommandRepository.DeleteProductByID(id); err != nil {
		return err
	}

	return nil
}
