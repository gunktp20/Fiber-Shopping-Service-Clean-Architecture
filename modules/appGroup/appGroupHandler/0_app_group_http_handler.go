package handler

import (
	"github.com/gofiber/fiber/v2"
	appGroupUsecase "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupUsecase"
)

type (
	AppGroupHttpHandlerService interface {
		CreateAppGroup(c *fiber.Ctx) error
	}

	appGroupHttpHandler struct {
		appGroupUsecase appGroupUsecase.AppGroupUsecaseService
	}
)

func NewAppGroupHttpHandler(usecase appGroupUsecase.AppGroupUsecaseService) AppGroupHttpHandlerService {
	return &appGroupHttpHandler{appGroupUsecase: usecase}
}
