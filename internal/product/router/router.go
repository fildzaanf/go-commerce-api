package router

import (
	"go-commerce-api/internal/product/handler"
	"go-commerce-api/internal/product/service"
	"go-commerce-api/internal/product/repository"
	
	"go-commerce-api/pkg/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductRouter(product *echo.Group, db *gorm.DB) {
	productQueryRepository := repository.NewProductQueryRepository(db)
	productCommandRepository := repository.NewProductCommandRepository(db)

	productQueryService := service.NewProductQueryService(productQueryRepository, productCommandRepository)
	productCommandService := service.NewProductCommandService(productCommandRepository, productQueryRepository)

	productHandler := handler.NewProductHandler(productCommandService, productQueryService)

	product.POST("", productHandler.CreateProduct, middleware.JWTMiddleware())
	product.PUT("/:id", productHandler.UpdateProductByID, middleware.JWTMiddleware())
	product.DELETE("/:id", productHandler.DeleteProductByID, middleware.JWTMiddleware())
	product.GET("/:id", productHandler.GetProductByID)
	product.GET("", productHandler.GetAllProducts)
}
