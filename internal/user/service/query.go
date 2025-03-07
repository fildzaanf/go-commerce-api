package service

import (
	"errors"
	"go-commerce-api/internal/user/domain"
	"go-commerce-api/internal/user/repository"
	"go-commerce-api/pkg/constant"
)

type userQueryService struct {
	userCommandRepository repository.UserCommandRepositoryInterface
	userQueryRepository   repository.UserQueryRepositoryInterface
}

func NewUserQueryService(ucr repository.UserCommandRepositoryInterface, uqr repository.UserQueryRepositoryInterface) UserQueryServiceInterface {
	return &userQueryService{
		userCommandRepository: ucr,
		userQueryRepository:   uqr,
	}
}

func (uqs *userQueryService) GetUserByID(id string) (domain.User, error) {
	if id == "" {
		return domain.User{}, errors.New(constant.ERROR_ID_INVALID)
	}

	userDomain, errGetID := uqs.userQueryRepository.GetUserByID(id)
	if errGetID != nil {
		return domain.User{}, errors.New(constant.ERROR_DATA_EMPTY)
	}

	return userDomain, nil
}
