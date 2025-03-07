package handler

import (
	"go-commerce-api/internal/user/dto"
	"go-commerce-api/internal/user/service"
	"go-commerce-api/pkg/constant"
	"go-commerce-api/pkg/middleware"
	"go-commerce-api/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userCommandService service.UserCommandServiceInterface
	userQueryService   service.UserQueryServiceInterface
}

func NewUserHandler(ucs service.UserCommandServiceInterface, uqs service.UserQueryServiceInterface) *userHandler {
	return &userHandler{
		userCommandService: ucs,
		userQueryService:   uqs,
	}
}

// command
func (uh *userHandler) RegisterUser(c echo.Context) error {
	var userRequest dto.UserRegisterRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	userDomain := dto.UserRegisterRequestToDomain(userRequest)

	registeredUser, err := uh.userCommandService.RegisterUser(userDomain)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	userResponse := dto.UserRegisterDomainToResponse(registeredUser)

	return c.JSON(http.StatusCreated, response.SuccessResponse(constant.SUCCESS_REGISTER, userResponse))
}

func (uh *userHandler) LoginUser(c echo.Context) error {
	var userRequest dto.UserLoginRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	loginUser, token, err := uh.userCommandService.LoginUser(userRequest.Email, userRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(err.Error()))
	}

	userResponse := dto.UserDomainToLoginResponse(loginUser, token)

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_LOGIN, userResponse))
}

// query
func (uh *userHandler) GetUserByID(c echo.Context) error {
	userIDParam := c.Param("id")
	if userIDParam == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("user id is required"))
	}

	userID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse("unauthorized access"))
	}

	if role != constant.USER || role != constant.SELLER || role != constant.BUYER  {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}
	
	if userIDParam != userID {
		return c.JSON(http.StatusForbidden, response.ErrorResponse("forbidden access"))
	}

	user, err := uh.userQueryService.GetUserByID(userIDParam)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse("user not found"))
	}

	userResponse := dto.UserDomainToResponse(user)

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, userResponse))
}
