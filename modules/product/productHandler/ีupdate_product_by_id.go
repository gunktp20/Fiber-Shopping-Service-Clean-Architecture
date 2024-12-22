package handler

import (
	"net/http"
	productDto "shopping-service-be/modules/product/productDto"
	"shopping-service-be/pkg/request"
	"shopping-service-be/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func (h *productHttpHandler) UpdateProductById(c *fiber.Ctx) error {

	contentType := c.Get("Content-Type")
	if contentType != "application/json" {
		return response.ErrResponse(c, http.StatusUnsupportedMediaType, "the content-type is not of type json")
	}

	selectedProductId := c.Params("product_id")

	wrapper := request.ContextWrapper(c)
	req := new(productDto.UpdateProductReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.productUsecase.UpdateProductById(selectedProductId, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
