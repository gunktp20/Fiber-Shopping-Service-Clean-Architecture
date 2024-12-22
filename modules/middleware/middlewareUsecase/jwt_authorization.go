package usecase

import (
	"shopping-service-be/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func (u *middlewareUsecase) JwtAuthorization(c fiber.Ctx, cfg *config.Config, accessToken string) (fiber.Ctx, error) {
	return c, nil
}
