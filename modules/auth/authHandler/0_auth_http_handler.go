package handler

import (
	authUsecase "shopping-service-be/modules/auth/authUsecase"

	"github.com/gofiber/fiber/v2"
)

type (
	AuthHttpHandlerService interface {
		Authenticate(c *fiber.Ctx) error
	}

	authHttpHandler struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(usecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{authUsecase: usecase}
}
