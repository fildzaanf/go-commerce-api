package service

import (
	"errors"
	"go-commerce-api/internal/product/domain"
	"go-commerce-api/internal/product/repository"
	"go-commerce-api/pkg/constant"
)

type productQueryService struct {
	productQueryRepository repository.ProductQueryRepositoryInterface
	productCommandRepository repository.ProductCommandRepositoryInterface
}

func NewProductQueryService(pqr repository.ProductQueryRepositoryInterface, pcr repository.ProductCommandRepositoryInterface) ProductQueryServiceInterface {
	return &productQueryService{
		productQueryRepository: pqr,
		productCommandRepository: pcr,
	}
}

func (pqs *productQueryService) GetProductByID(id string) (domain.Product, error) {
	if id == "" {
		return domain.Product{}, errors.New(constant.ERROR_ID_INVALID)
	}

	product, err := pqs.productQueryRepository.GetProductByID(id)
	if err != nil {
		return domain.Product{}, errors.New("product not found")
	}

	return product, nil
}

func (pqs *productQueryService) GetAllProducts() ([]domain.Product, error) {
	products, err := pqs.productQueryRepository.GetAllProducts()

	if err != nil {
		return nil, errors.New(constant.ERROR_DATA_EMPTY)
	}

	return products, nil
}
