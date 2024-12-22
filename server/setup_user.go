package server

import (
	userHandler "shopping-service-be/modules/user/userHandler"
	userRepository "shopping-service-be/modules/user/userRepository"
	userRouter "shopping-service-be/modules/user/userRouter"
	userUsecase "shopping-service-be/modules/user/userUsecase"

	"github.com/gofiber/fiber/v2"
)

func (s *fiberServer) initializeUserHttpHandler(api fiber.Router) {
	// ? Initialize all layers
	userPostgresRepository := userRepository.NewUserPostgresRepository(s.db.GetDb())

	userUsecase := userUsecase.NewUserUsecase(userPostgresRepository)
	userHttpHandler := userHandler.NewUserHttpHandler(userUsecase)

	// Routers
	userRouter.SetupUserRoutes(api, userHttpHandler)
}
