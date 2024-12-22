package handler

import (
	productUsecase "shopping-service-be/modules/product/productUsecase"

	"github.com/gofiber/fiber/v2"
)

type (
	ProductHttpHandlerService interface {
		CreateProduct(c *fiber.Ctx) error
		GetAllProducts(c *fiber.Ctx) error
		UpdateProductById(c *fiber.Ctx) error
		DeleteProductById(c *fiber.Ctx) error

		CreateProductFormData(c *fiber.Ctx) error
		CreateProductJson(c *fiber.Ctx) error
		CreateProductUrlencoded(c *fiber.Ctx) error

		UploadProductImage(c *fiber.Ctx) error
	}

	productHttpHandler struct {
		productUsecase productUsecase.ProductUsecaseService
	}
)

func NewProductHttpHandler(usecase productUsecase.ProductUsecaseService) ProductHttpHandlerService {
	return &productHttpHandler{productUsecase: usecase}
}
