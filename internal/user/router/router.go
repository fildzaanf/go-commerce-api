package router

import (
	"go-commerce-api/internal/user/handler"
	"go-commerce-api/internal/user/repository"
	"go-commerce-api/internal/user/service"
	"go-commerce-api/pkg/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRouter(user *echo.Group, db *gorm.DB) {
	userQueryRepository := repository.NewUserQueryRepository(db)
	userCommandRepository := repository.NewUserCommandRepository(db)

	userQueryService := service.NewUserQueryService(userCommandRepository, userQueryRepository)
	userCommandService := service.NewUserCommandService(userCommandRepository, userQueryRepository)

	userHandler := handler.NewUserHandler(userCommandService, userQueryService)

	user.POST("/register", userHandler.RegisterUser)
	user.POST("/login", userHandler.LoginUser)
	user.GET("/:user_id", userHandler.GetUserByID, middleware.JWTMiddleware())
}
