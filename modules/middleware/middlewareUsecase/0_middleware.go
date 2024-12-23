package usecase

import (
	"shopping-service-be/pkg/config"

	"github.com/golang-jwt/jwt/v4"
)

type (
	MiddlewareUsecaseService interface {
		JwtAuthorization(tokenString string) (*jwt.Token, error)
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
