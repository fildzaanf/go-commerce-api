package domain

import (
	"go-commerce-api/internal/user/entity"
	"time"
)

type User struct {
	ID              string
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func UserDomainToEntity(userDomain User) entity.User {
	return entity.User{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Password:  userDomain.Password,
		Role:      userDomain.Role,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}

func UserEntityToDomain(userEntity entity.User) User {
	return User{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		Password:  userEntity.Password,
		Role:      userEntity.Role,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}
}
