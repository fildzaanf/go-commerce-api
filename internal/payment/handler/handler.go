package handler

import (
	"encoding/json"
	"fmt"
	"go-commerce-api/internal/payment/dto"
	"go-commerce-api/internal/payment/service"
	"go-commerce-api/pkg/constant"
	"go-commerce-api/pkg/middleware"
	"go-commerce-api/pkg/response"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type paymentHandler struct {
	paymentCommandService service.PaymentCommandServiceInterface
	paymentQueryService   service.PaymentQueryServiceInterface
}

func NewPaymentHandler(pcs service.PaymentCommandServiceInterface, pqs service.PaymentQueryServiceInterface) *paymentHandler {
	return &paymentHandler{
		paymentCommandService: pcs,
		paymentQueryService:   pqs,
	}
}

func (ph *paymentHandler) CreatePayment(c echo.Context) error {
	userID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse("unauthorized access"))
	}

	if role != constant.BUYER {
		return c.JSON(http.StatusForbidden, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	var paymentRequest dto.CreatePaymentRequest
	if err := c.Bind(&paymentRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	paymentDomain := dto.CreatePaymentRequestToDomain(paymentRequest)

	createdPayment, err := ph.paymentCommandService.CreatePayment(paymentDomain, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	paymentResponse := dto.PaymentDomainToResponse(createdPayment)

	return c.JSON(http.StatusCreated, response.SuccessResponse("payment created successfully", paymentResponse))
}

func (ph *paymentHandler) MidtransWebhook(c echo.Context) error {
	var notification map[string]interface{}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("error reading request body:", err)
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("invalid request body"))
	}

	if err := json.Unmarshal(body, &notification); err != nil {
		fmt.Println("error decoding json:", err)
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("invalid json format"))
	}

	orderID, ok := notification["order_id"].(string)
	if !ok {
		fmt.Println("missing order_id in request")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("missing order_id"))
	}

	transactionStatus, ok := notification["transaction_status"].(string)
	if !ok {
		fmt.Println("missing transaction_status in request")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("missing transaction_status"))
	}

	var updatedStatus string
	switch transactionStatus {
	case "settlement":
		updatedStatus = "success"
	case "expire":
		updatedStatus = "expired"
	case "cancel":
		updatedStatus = "cancel"
	case "deny":
		updatedStatus = "deny"
	default:
		updatedStatus = "pending"
	}

	err = ph.paymentCommandService.UpdatePaymentStatusByID(orderID, updatedStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("failed to update payment status"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse("payment status updated", nil))
}

func (ph *paymentHandler) GetPaymentByID(c echo.Context) error {
	userID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse("unauthorized access"))
	}

	if role != constant.BUYER && role != constant.SELLER {
		return c.JSON(http.StatusForbidden, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}	

	paymentID := c.Param("id")
	if paymentID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("payment id is required"))
	}

	payment, err := ph.paymentQueryService.GetPaymentByID(paymentID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse("Payment not found"))
	}

	if payment.UserID != userID {
		return c.JSON(http.StatusForbidden, response.ErrorResponse("forbidden access"))
	}

	paymentResponse := dto.PaymentDomainToResponse(payment)

	return c.JSON(http.StatusOK, response.SuccessResponse("Payment retrieved successfully", paymentResponse))
}

func (ph *paymentHandler) GetAllPayments(c echo.Context) error {
	userID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse("unauthorized access"))
	}

	if role != constant.BUYER && role != constant.SELLER {
		return c.JSON(http.StatusForbidden, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}	

	payments, err := ph.paymentQueryService.GetAllPayments(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to retrieve payments"))
	}

	paymentResponses := dto.ListPaymentDomainToResponse(payments)
	
	return c.JSON(http.StatusOK, response.SuccessResponse("Payments retrieved successfully", paymentResponses))
}
