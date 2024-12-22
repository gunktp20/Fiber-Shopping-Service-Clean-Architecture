package handler

import (
	"net/http"
	dto "shopping-service-be/modules/product/productDto"
	helpers "shopping-service-be/pkg/helpers"
	"shopping-service-be/pkg/request"
	"shopping-service-be/pkg/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (h *productHttpHandler) CreateProductFormData(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")
	if !strings.Contains(contentType, "multipart/form-data") {
		return response.ErrResponse(c, http.StatusUnsupportedMediaType, "the content-type is not of type multipart/form-data")
	}

	fileHeader, err := c.FormFile("image")
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "File is required")
	}

	fileBytes, err := helpers.ConvertMultipartFileToBytes(fileHeader)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Failed to convert multipart file to bytes")
	}

	allowedTypes := []string{"image/jpeg", "image/png", "image/gif"}
	maxFileSize := int64(5 * 1024 * 1024)

	if err := helpers.ValidateFile(fileBytes, allowedTypes, maxFileSize); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	wrapper := request.ContextWrapper(c)
	body := new(dto.CreateProductReq)

	if err := wrapper.Bind(body); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.productUsecase.CreateProductFormData(body, fileBytes)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
