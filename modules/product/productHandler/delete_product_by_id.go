package handler

import (
	"net/http"
	"shopping-service-be/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func (h *productHttpHandler) DeleteProductById(c *fiber.Ctx) error {

	productId := c.Params("product_id")

	err := h.productUsecase.DeleteProductById(productId)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, fiber.Map{
		"message": "Deleted 1 product",
	})
}
