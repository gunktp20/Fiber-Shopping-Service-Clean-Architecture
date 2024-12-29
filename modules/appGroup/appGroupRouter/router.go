package router

import (
	"github.com/gofiber/fiber/v2"
	appGroupHandler "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupHandler"
)

func SetAppGroupRoutes(api fiber.Router, appGroupHttpHandler appGroupHandler.AppGroupHttpHandlerService) {
	routes := api.Group("/app-group")

	routes.Post("/", appGroupHttpHandler.CreateAppGroup)

}
