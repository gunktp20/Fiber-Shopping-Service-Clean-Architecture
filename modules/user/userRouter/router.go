package router

import (
	handler "shopping-service-be/modules/user/userHandler"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router, productHttpHandler handler.UserHttpHandlerService) {
	routes := api.Group("/user")

	routes.Post("/", productHttpHandler.CreateUser)

}
