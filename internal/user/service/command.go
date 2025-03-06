package service

import (
	"errors"
	"go-commerce-api/internal/user/domain"
	"go-commerce-api/internal/user/repository"
	"go-commerce-api/pkg/constant"
	"go-commerce-api/pkg/crypto"
	"go-commerce-api/pkg/middleware"
	"go-commerce-api/pkg/validator"
)

type userCommandService struct {
	userCommandRepository repository.UserCommandRepositoryInterface
	userQueryRepository   repository.UserQueryRepositoryInterface
}

func NewUserCommandService(ucr repository.UserCommandRepositoryInterface, uqr repository.UserQueryRepositoryInterface) UserCommandServiceInterface {
	return &userCommandService{
		userCommandRepository: ucr,
		userQueryRepository:   uqr,
	}
}

func (ucs *userCommandService) RegisterUser(user domain.User) (domain.User, error) {

	errEmpty := validator.IsDataEmpty(
		[]string{"name", "email", "password", "confirm_password"},
		user.Name, user.Email, user.Password, user.ConfirmPassword,
	)
	if errEmpty != nil {
		return domain.User{}, errEmpty
	}

	errEmailValid := validator.IsEmailValid(user.Email)
	if errEmailValid != nil {
		return domain.User{}, errEmailValid
	}

	errLength := validator.IsMinLengthValid(8, map[string]string{"password": user.Password})
	if errLength != nil {
		return domain.User{}, errLength
	}

	_, errGetEmail := ucs.userQueryRepository.GetUserByEmail(user.Email)
	if errGetEmail == nil {
		return domain.User{}, errors.New(constant.ERROR_EMAIL_EXIST)
	}

	if user.Password != user.ConfirmPassword {
		return domain.User{}, errors.New(constant.ERROR_PASSWORD_CONFIRM)
	}

	hashedPassword, errHash := crypto.HashPassword(user.Password)
	if errHash != nil {
		return domain.User{}, errors.New(constant.ERROR_PASSWORD_HASH)
	}

	user.Password = hashedPassword

	userEntity, errRegister := ucs.userCommandRepository.RegisterUser(user)
	if errRegister != nil {
		return domain.User{}, errRegister
	}

	return userEntity, nil
}

func (ucs *userCommandService) LoginUser(email, password string) (domain.User, string, error) {
	errEmpty := validator.IsDataEmpty([]string{"email", "password"}, email, password)
	if errEmpty != nil {
		return domain.User{}, "", errEmpty
	}

	errEmailValid := validator.IsEmailValid(email)
	if errEmailValid != nil {
		return domain.User{}, "", errEmailValid
	}

	userDomain, errGetEmail := ucs.userQueryRepository.GetUserByEmail(email)
	if errGetEmail != nil {
		return domain.User{}, "", errors.New(constant.ERROR_EMAIL_UNREGISTERED)
	}

	comparePassword := crypto.ComparePassword(userDomain.Password, password)
	if comparePassword != nil {
		return domain.User{}, "", errors.New(constant.ERROR_LOGIN)
	}

	token, errCreate := middleware.GenerateToken(userDomain.ID, userDomain.Role)
	if errCreate != nil {
		return domain.User{}, "", errors.New(constant.ERROR_TOKEN_GENERATE)
	}

	return userDomain, token, nil
}
