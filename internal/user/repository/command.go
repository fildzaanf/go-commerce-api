package repository

import (
	"errors"
	"go-commerce-api/internal/user/domain"
	"go-commerce-api/internal/user/entity"
	"go-commerce-api/pkg/constant"
	"go-commerce-api/pkg/crypto"

	"gorm.io/gorm"
)

type userCommandRepository struct {
	db *gorm.DB
}

func NewUserCommandRepository(db *gorm.DB) UserCommandRepositoryInterface {
	return &userCommandRepository{
		db: db,
	}
}

func (ucr *userCommandRepository) RegisterUser(user domain.User) (domain.User, error) {
	tx := ucr.db.Begin()
	if tx.Error != nil {
		return domain.User{}, tx.Error
	}

	userEntity := domain.UserDomainToEntity(user)

	if err := tx.Create(&userEntity).Error; err != nil {
		tx.Rollback()
		return domain.User{}, err
	}

	userDomain := domain.UserEntityToDomain(userEntity)

	if err := tx.Commit().Error; err != nil {
		return domain.User{}, err
	}

	return userDomain, nil
}

func (ucr *userCommandRepository) LoginUser(email, password string) (domain.User, error) {
	tx := ucr.db.Begin()

	if tx.Error != nil {
		return domain.User{}, tx.Error
	}

	userEntity := entity.User{}

	result := tx.Where("email = ?", email).First(&userEntity)
	if result.Error != nil {
		tx.Rollback()
		return domain.User{}, result.Error
	}

	if errComparePass := crypto.ComparePassword(userEntity.Password, password); errComparePass != nil {
		tx.Rollback()
		return domain.User{}, errors.New(constant.ERROR_PASSWORD_INVALID)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.User{}, err
	}

	userDomain := domain.UserEntityToDomain(userEntity)

	return userDomain, nil
}
