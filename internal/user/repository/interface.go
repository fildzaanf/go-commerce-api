package repository

import (
	"go-commerce-api/internal/user/domain"
)

type UserCommandRepositoryInterface interface {
	RegisterUser(user domain.User) (domain.User, error)
	LoginUser(email, password string) (domain.User, error)
}

type UserQueryRepositoryInterface interface {
	GetUserByID(id string) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
}
