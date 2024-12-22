package handler

import (
	userUsecase "shopping-service-be/modules/user/userUsecase"

	"github.com/gofiber/fiber/v2"
)

type (
	UserHttpHandlerService interface {
		CreateUser(c *fiber.Ctx) error
	}

	userHttpHandler struct {
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserHttpHandler(usecase userUsecase.UserUsecaseService) UserHttpHandlerService {
	return &userHttpHandler{userUsecase: usecase}
}
