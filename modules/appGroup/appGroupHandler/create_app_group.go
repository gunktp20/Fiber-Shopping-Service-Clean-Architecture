package handler

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	appGroupDto "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupDto"
	"github.com/gunktp20/digital-hubx-be/pkg/response"
	"github.com/gunktp20/digital-hubx-be/pkg/utils"
)

func (h *appGroupHttpHandler) CreateAppGroup(c *fiber.Ctx) error {
	var body appGroupDto.CreateAppGroupReq
	if err := c.BodyParser(&body); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "The input data is invalid", nil)
	}

	if err := validator.New().Struct(&body); err != nil {
		validationErrors := utils.TranslateValidationError(err.(validator.ValidationErrors))
		return response.ErrResponse(c, http.StatusBadRequest, "The input data is invalid", &validationErrors)
	}

	res, err := h.appGroupUsecase.CreateAppGroup(&body)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error(), nil)
	}

	return response.SuccessResponse(c, http.StatusOK, res)

}
