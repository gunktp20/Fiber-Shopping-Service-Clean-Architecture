package handler

import (
	"net/http"
	productDto "shopping-service-be/modules/product/productDto"
	"shopping-service-be/pkg/request"
	"shopping-service-be/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func (h *productHttpHandler) CreateProductJson(c *fiber.Ctx) error {

	contentType := c.Get("Content-Type")
	if contentType != "application/json" {
		return response.ErrResponse(c, http.StatusUnsupportedMediaType, "the content-type is not of type json")
	}

	wrapper := request.ContextWrapper(c)
	req := new(productDto.CreateProductReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.productUsecase.CreateProductJson(req)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
