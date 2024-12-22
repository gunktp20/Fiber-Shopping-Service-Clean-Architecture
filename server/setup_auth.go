package server

import (
	authHandler "shopping-service-be/modules/auth/authHandler"
	authRouter "shopping-service-be/modules/auth/authRouter"
	authUsecase "shopping-service-be/modules/auth/authUsecase"
	"shopping-service-be/pkg/config"

	userRepository "shopping-service-be/modules/user/userRepository"

	"github.com/gofiber/fiber/v2"
)

func (s *fiberServer) initializeAuthHttpHandler(api fiber.Router, conf *config.Config) {
	// ? Initialize all layers
	userPostgresRepository := userRepository.NewUserPostgresRepository(s.db.GetDb())

	authUsecase := authUsecase.NewAuthUsecase(userPostgresRepository, conf)
	authHttpHandler := authHandler.NewAuthHttpHandler(authUsecase)

	// Routers
	authRouter.SetupAuthRoutes(api, authHttpHandler)
}
