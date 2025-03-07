package handler

import "github.com/labstack/echo/v4"

type UserHandlerInterface interface {
	// command
	RegisterUser(c echo.Context) error
	LoginUser(c echo.Context) error

	// query
	GetUserByID(c echo.Context) error
}
