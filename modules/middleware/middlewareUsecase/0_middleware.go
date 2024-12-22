package usecase

import (
	"shopping-service-be/pkg/config"

	"github.com/gofiber/fiber/v2"
)

type (
	MiddlewareUsecaseService interface {
		JwtAuthorization(c fiber.Ctx, cfg *config.Config, accessToken string) (fiber.Ctx, error)
	}

	middlewareUsecase struct {
		cfg *config.Config
	}
)

func NewMiddlewareUsecase(cfg *config.Config) MiddlewareUsecaseService {
	return &middlewareUsecase{
		cfg: cfg,
	}
}
