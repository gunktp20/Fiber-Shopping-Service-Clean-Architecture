package handler

import (
	"net/http"
	userDto "shopping-service-be/modules/user/userDto"
	"shopping-service-be/pkg/request"
	"shopping-service-be/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func (h *userHttpHandler) CreateUser(c *fiber.Ctx) error {
	wrapper := request.ContextWrapper(c)
	req := new(userDto.CreateUserReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.CreateUser(req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
