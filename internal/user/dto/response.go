package dto

import "go-commerce-api/internal/user/domain"

type UserRegisterResponse struct {
	ID    string `json:"id"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

type UserLoginResponse struct {
	ID    string `json:"id"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func UserRegisterDomainToResponse(user domain.User) UserRegisterResponse {
	return UserRegisterResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
}

func UserDomainToLoginResponse(user domain.User, token string) UserLoginResponse {
	return UserLoginResponse{
		ID:    user.ID,
		Role:  user.Role,
		Token: token,
	}
}

func UserDomainToResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}
