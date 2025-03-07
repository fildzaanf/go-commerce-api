package router

import (
	"go-commerce-api/internal/payment/handler"
	"go-commerce-api/internal/payment/repository"
	"go-commerce-api/internal/payment/service"
	"go-commerce-api/pkg/middleware"
	repositoryProduct "go-commerce-api/internal/product/repository"
	repositoryUser "go-commerce-api/internal/user/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PaymentRouter(payment *echo.Group, db *gorm.DB) {
	paymentQueryRepository := repository.NewPaymentQueryRepository(db)
    paymentCommandRepository := repository.NewPaymentCommandRepository(db)
    productQueryRepository := repositoryProduct.NewProductQueryRepository(db)
    productCommandRepository := repositoryProduct.NewProductCommandRepository(db)
    userQueryRepository := repositoryUser.NewUserQueryRepository(db)

    paymentQueryService := service.NewPaymentQueryService(paymentQueryRepository, paymentCommandRepository)
    paymentCommandService := service.NewPaymentCommandService(paymentCommandRepository, paymentQueryRepository, productQueryRepository, productCommandRepository, userQueryRepository,)


	paymentHandler := handler.NewPaymentHandler(paymentCommandService, paymentQueryService)

	payment.POST("", paymentHandler.CreatePayment, middleware.JWTMiddleware())
	payment.GET("/:id", paymentHandler.GetPaymentByID, middleware.JWTMiddleware())
	payment.GET("", paymentHandler.GetAllPayments, middleware.JWTMiddleware())
	payment.POST("/midtrans/webhook", paymentHandler.MidtransWebhook)
}
