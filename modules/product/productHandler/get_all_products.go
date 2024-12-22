package handler

import (
	"net/http"
	"shopping-service-be/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func (h *productHttpHandler) GetAllProducts(c *fiber.Ctx) error {

	searchKey := c.Query("search", "")
	res, err := h.productUsecase.GetAllProducts(searchKey)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
