package handler

import appGroupUsecase "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupUsecase"

type (
	AppGroupHttpHandlerService interface {
	}

	appGroupHttpHandler struct {
		appGroupUsecase appGroupUsecase.AppGroupUsecaseService
	}
)

func NewAppGroupHttpHandler(usecase appGroupUsecase.AppGroupUsecaseService) AppGroupHttpHandlerService {
	return &appGroupHttpHandler{appGroupUsecase: usecase}
}
