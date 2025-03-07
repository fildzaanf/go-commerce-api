package handler

import "github.com/labstack/echo/v4"

type ProductHandlerInterface interface {
	// command
	CreatePayment(c echo.Context) error
	MidtransWebhook(c echo.Context) error

	// query
	GetPaymentByID(c echo.Context) error
	GetAllPayments(c echo.Context) error
}
