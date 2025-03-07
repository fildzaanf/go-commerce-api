package service

import (
	"errors"
	"go-commerce-api/infrastructure/config"
	"go-commerce-api/internal/payment/domain"
	"go-commerce-api/internal/payment/repository"
	repositoryProduct "go-commerce-api/internal/product/repository"
	repositoryUser "go-commerce-api/internal/user/repository"
	"log"

	"go-commerce-api/pkg/constant"
	"go-commerce-api/pkg/email/mailer"
	"go-commerce-api/pkg/generator"
	"go-commerce-api/pkg/validator"
	"time"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/shopspring/decimal"
)

type paymentCommandService struct {
	paymentCommandRepository repository.PaymentCommandRepositoryInterface
	paymentQueryRepository   repository.PaymentQueryRepositoryInterface
	productQueryRepository   repositoryProduct.ProductQueryRepositoryInterface
	productCommandRepository repositoryProduct.ProductCommandRepositoryInterface
	userQueryRepository      repositoryUser.UserQueryRepositoryInterface
}

func NewPaymentCommandService(pycr repository.PaymentCommandRepositoryInterface, pyqr repository.PaymentQueryRepositoryInterface, pdqr repositoryProduct.ProductQueryRepositoryInterface, pdcr repositoryProduct.ProductCommandRepositoryInterface, uqr repositoryUser.UserQueryRepositoryInterface) PaymentCommandServiceInterface {
	return &paymentCommandService{
		paymentCommandRepository: pycr,
		paymentQueryRepository:   pyqr,
		productQueryRepository:   pdqr,
		productCommandRepository: pdcr,
		userQueryRepository:      uqr,
	}
}

func (pcs *paymentCommandService) CreatePayment(payment domain.Payment, userID string) (domain.Payment, error) {
	errEmpty := validator.IsDataEmpty([]string{"product_id", "quantity"}, payment.ProductID, payment.Quantity)
	if errEmpty != nil {
		return domain.Payment{}, errEmpty
	}

	product, err := pcs.productQueryRepository.GetProductByID(payment.ProductID)
	if err != nil {
		return domain.Payment{}, errors.New("product not found")
	}

	if product.Stock < payment.Quantity {
		return domain.Payment{}, errors.New("insufficient stock")
	}

	product.Stock -= payment.Quantity
	errUpdateStock := pcs.productCommandRepository.UpdateProductStockByID(product.ID, product.Stock)
	if errUpdateStock != nil {
		return domain.Payment{}, errors.New("failed to update product stock")
	}

	quantityDecimal := decimal.NewFromInt(int64(payment.Quantity))
	payment.TotalAmount = product.Price.Mul(quantityDecimal)

	user, err := pcs.userQueryRepository.GetUserByID(userID)
	if err != nil {
		return domain.Payment{}, errors.New("User not found")
	}

	if user.ID != userID {
		return domain.Payment{}, errors.New(constant.ERROR_ROLE_ACCESS)
	}

	payment.PaymentCode = generator.GeneratePaymentCode()

	cfg, err := config.LoadConfig()
	if err != nil {
		return domain.Payment{}, errors.New("failed to load configuration")
	}

	midtransClient := snap.Client{}
	midtransClient.New(cfg.MIDTRANS.MIDTRANS_SERVER_KEY, midtrans.Sandbox)

	payment.ID = uuid.New().String()

	snapRequest := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.ID,
			GrossAmt: payment.TotalAmount.IntPart(),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	snapResponse, errSnap := midtransClient.CreateTransaction(snapRequest)
	if errSnap != nil {
		product.Stock += payment.Quantity
		_ = pcs.productCommandRepository.UpdateProductStockByID(product.ID, product.Stock)
		return domain.Payment{}, errSnap
	}

	payment.UserID = userID
	payment.Status = "pending"
	payment.PaymentURL = snapResponse.RedirectURL
	payment.Token = snapResponse.Token
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()

	paymentEntity, errCreate := pcs.paymentCommandRepository.CreatePayment(payment)
	if errCreate != nil {
		product.Stock += payment.Quantity
		_ = pcs.productCommandRepository.UpdateProductStockByID(product.ID, product.Stock)
		return domain.Payment{}, errCreate
	}

	return paymentEntity, nil
}

func (pcs *paymentCommandService) UpdatePaymentStatusByID(id, status string) error {
	payment, err := pcs.paymentQueryRepository.GetPaymentByID(id)
	if err != nil {
		return errors.New("payment not found")
	}

	if payment.ID != id {
		return errors.New(constant.ERROR_ID_NOT_FOUND)
	}

	user, err := pcs.userQueryRepository.GetUserByID(payment.UserID)
	if err != nil {
		return errors.New("failed to retrieve user information")
	}

	product, err := pcs.productQueryRepository.GetProductByID(payment.ProductID)
	if err != nil {
		return errors.New("failed to retrieve product information")
	}

	prevStatus := payment.Status

	switch status {
	case "settlement":
		payment.Status = "success"
	case "expired":
		payment.Status = "expire"
	case "cancel":
		payment.Status = "cancel"
	case "deny":
		payment.Status = "deny"
	default:
		payment.Status = status
	}

	log.Printf("received payment status from Midtrans: %s", status)

	if prevStatus == "pending" && (status == "expired" || status == "cancel" || status == "deny") {
		product.Stock += payment.Quantity
		err := pcs.productCommandRepository.UpdateProductStockByID(product.ID, product.Stock)
		if err != nil {
			return errors.New("failed to restore product stock")
		}
	}

	err = pcs.paymentCommandRepository.UpdatePaymentStatusByID(id, payment.Status)
	if err != nil {
		return errors.New("failed to update payment status")
	}

	mailer.SendEmailNotificationPayment(
		user.Name,
		user.Email,
		payment.PaymentCode,
		product.Name,
		product.Price,
		payment.Quantity,
		payment.TotalAmount,
		payment.Status,
		payment.UpdatedAt,
	)

	return nil
}
