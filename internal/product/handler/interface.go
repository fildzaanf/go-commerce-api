package handler

import "github.com/labstack/echo/v4"

type ProductHandlerInterface interface {
	// command
	CreateProduct(c echo.Context) error
	UpdateProductByID(c echo.Context) error
	DeleteProductByID(c echo.Context) error

	// query
	GetProductByID(c echo.Context) error
	GetAllProducts(c echo.Context) error
}
