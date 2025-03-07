package handler

import (
	"go-commerce-api/internal/product/dto"
	"go-commerce-api/internal/product/service"
	"go-commerce-api/pkg/cloud"
	"go-commerce-api/pkg/constant"
	"go-commerce-api/pkg/middleware"
	"go-commerce-api/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productCommandService service.ProductCommandServiceInterface
	productQueryService   service.ProductQueryServiceInterface
}

func NewProductHandler(pcs service.ProductCommandServiceInterface, pqs service.ProductQueryServiceInterface) *productHandler {
	return &productHandler{
		productCommandService: pcs,
		productQueryService:   pqs,
	}
}

// command
func (ph *productHandler) CreateProduct(c echo.Context) error {
	userID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse("unauthorized access"))
	}

	if role != constant.SELLER {
		return c.JSON(http.StatusForbidden, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	var productRequest dto.CreateProductRequest
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	productDomain := dto.CreateProductRequestToDomain(productRequest, userID)

	image, errImage := c.FormFile("image_url")
	if errImage != nil && errImage != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(constant.ERROR_UPLOAD_IMAGE))
	}

	if image != nil {
		imageURL, errUpload := cloud.UploadImageToS3(image)
		if errUpload != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse(constant.ERROR_UPLOAD_IMAGE_S3))
		}
		productRequest.ImageURL = imageURL
	}

	createdProduct, err := ph.productCommandService.CreateProduct(productDomain, image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	productResponse := dto.ProductDomainToResponse(createdProduct)

	return c.JSON(http.StatusCreated, response.SuccessResponse(constant.SUCCESS_CREATED, productResponse))
}

func (ph *productHandler) UpdateProductByID(c echo.Context) error {
	userID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse("unauthorized access"))
	}

	productID := c.Param("id")
	if productID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("product id is required"))
	}

	product, err := ph.productQueryService.GetProductByID(productID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse("product not found"))
	}

	if product.UserID != userID {
		return c.JSON(http.StatusForbidden, response.ErrorResponse("forbidden access"))
	}

	if role != constant.SELLER {
		return c.JSON(http.StatusForbidden, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	var productRequest dto.UpdateProductRequest
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	image, errImage := c.FormFile("image_url")
	if errImage != nil && errImage != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(constant.ERROR_UPLOAD_IMAGE))
	}

	if image != nil {
		imageURL, errUpload := cloud.UploadImageToS3(image)
		if errUpload != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse(constant.ERROR_UPLOAD_IMAGE_S3))
		}
		productRequest.ImageURL = imageURL
	}

	productDomain := dto.UpdateProductRequestToDomain(productRequest)

	updatedProduct, err := ph.productCommandService.UpdateProductByID(productID, productDomain, image)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	productResponse := dto.ProductDomainToResponse(updatedProduct)
	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_UPDATED, productResponse))
}

func (ph *productHandler) DeleteProductByID(c echo.Context) error {
	userID, role, errExtract := middleware.ExtractToken(c)
	if errExtract != nil {
		return c.JSON(http.StatusUnauthorized, response.ErrorResponse("unauthorized access"))
	}

	productID := c.Param("id")
	if productID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("product id is required"))
	}

	product, err := ph.productQueryService.GetProductByID(productID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse("product not found"))
	}

	if product.UserID != userID {
		return c.JSON(http.StatusForbidden, response.ErrorResponse("forbidden access"))
	}

	if role != constant.SELLER {
		return c.JSON(http.StatusForbidden, response.ErrorResponse(constant.ERROR_ROLE_ACCESS))
	}

	if err := ph.productCommandService.DeleteProductByID(productID); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse("product successfully deleted", nil))
}

// query
func (ph *productHandler) GetProductByID(c echo.Context) error {
	productID := c.Param("id")
	if productID == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("product id is required"))
	}

	product, err := ph.productQueryService.GetProductByID(productID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse("product not found"))
	}

	productResponse := dto.ProductDomainToResponse(product)
	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, productResponse))
}

func (ph *productHandler) GetAllProducts(c echo.Context) error {
	products, err := ph.productQueryService.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse("failed to retrieve products"))
	}

	productResponses := dto.ListProductDomainToResponse(products)

	return c.JSON(http.StatusOK, response.SuccessResponse(constant.SUCCESS_RETRIEVED, productResponses))
}
