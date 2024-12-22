package server

import (
	productHandler "shopping-service-be/modules/product/productHandler"
	productRepository "shopping-service-be/modules/product/productRepository"
	productRouter "shopping-service-be/modules/product/productRouter"
	productUsecase "shopping-service-be/modules/product/productUsecase"

	"github.com/gofiber/fiber/v2"
)

func (s *fiberServer) initializeProductHttpHandler(api fiber.Router) {
	// ? Initialize all layers
	productPostgresRepository := productRepository.NewProductPostgresRepository(s.db.GetDb())

	productUsecase := productUsecase.NewProductUsecase(productPostgresRepository)
	productHttpHandler := productHandler.NewProductHttpHandler(productUsecase)

	// Routers
	productRouter.SetupProductRoutes(api, productHttpHandler)
}
