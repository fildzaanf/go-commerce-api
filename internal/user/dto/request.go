package dto

import "go-commerce-api/internal/user/domain"

type UserRegisterRequest struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	Role            string `json:"role" form:"role"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func UserRegisterRequestToDomain(request UserRegisterRequest) domain.User {
	return domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Role:     request.Role,
		Password: request.Password,
	}
}

func UserLoginRequestToDomain(request UserLoginRequest) domain.User {
	return domain.User{
		Email:    request.Email,
		Password: request.Password,
	}
}
