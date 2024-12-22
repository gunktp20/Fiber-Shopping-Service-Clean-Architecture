package handler

import (
	middlewareUsecase "shopping-service-be/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface {
	}

	middlewareHandler struct {
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHttpHandler(usecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{middlewareUsecase: usecase}
}
