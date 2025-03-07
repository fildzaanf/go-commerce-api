package service

import (
	"errors"
	"go-commerce-api/internal/payment/domain"
	"go-commerce-api/internal/payment/repository"
	"go-commerce-api/pkg/constant"
)

type paymentQueryService struct {
	paymentCommandRepository repository.PaymentCommandRepositoryInterface
	paymentQueryRepository   repository.PaymentQueryRepositoryInterface
}

func NewPaymentQueryService(pqr repository.PaymentQueryRepositoryInterface, pcr repository.PaymentCommandRepositoryInterface) PaymentQueryServiceInterface {
	return &paymentQueryService{
		paymentQueryRepository:   pqr,
		paymentCommandRepository: pcr,
	}
}

func (pqs *paymentQueryService) GetPaymentByID(id string) (domain.Payment, error) {
	if id == "" {
		return domain.Payment{}, errors.New(constant.ERROR_ID_INVALID)
	}

	payment, err := pqs.paymentQueryRepository.GetPaymentByID(id)
	if err != nil {
		return domain.Payment{}, err
	}

	return payment, nil
}

func (pqs *paymentQueryService) GetAllPayments(userID string) ([]domain.Payment, error) {
	if userID == "" {
		return nil, errors.New(constant.ERROR_ID_INVALID)
	}

	payments, err := pqs.paymentQueryRepository.GetAllPayments(userID)
	if err != nil {
		return nil, errors.New(constant.ERROR_DATA_EMPTY)
	}

	return payments, nil
}
