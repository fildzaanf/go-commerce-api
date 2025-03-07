package service

import "go-commerce-api/internal/user/domain"

type UserCommandServiceInterface interface {
	RegisterUser(user domain.User) (domain.User, error)
	LoginUser(email, password string) (domain.User, string, error)
}

type UserQueryServiceInterface interface {
	GetUserByID(id string) (domain.User, error)
}
