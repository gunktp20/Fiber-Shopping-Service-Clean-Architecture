package handler

import (
	"net/http"
	"shopping-service-be/pkg/helpers"
	"shopping-service-be/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func (h *productHttpHandler) UploadProductImage(c *fiber.Ctx) error {

	fileBytes := c.Body()
	if len(fileBytes) == 0 {
		return response.ErrResponse(c, http.StatusBadRequest, "File is required")
	}

	allowedTypes := []string{"image/jpeg", "image/png", "image/gif"}
	maxFileSize := int64(5 * 1024 * 1024) // 5 MB

	if err := helpers.ValidateFile(fileBytes, allowedTypes, maxFileSize); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	res, err := h.productUsecase.UploadProductImage(fileBytes)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
