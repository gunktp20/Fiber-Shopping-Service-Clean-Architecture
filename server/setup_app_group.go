package server

import (
	"github.com/gofiber/fiber/v2"
	appGroupHandler "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupHandler"
	appGroupRepository "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupRepository"
	appGroupRouter "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupRouter"
	appGroupUsecase "github.com/gunktp20/digital-hubx-be/modules/appGroup/appGroupUsecase"
	"github.com/gunktp20/digital-hubx-be/pkg/config"
)

func (s *fiberServer) initializeAppGroupHttpHandler(api fiber.Router, conf *config.Config) {
	// ? Initialize all layers
	appGroupGormRepository := appGroupRepository.NewAppGroupGormRepository(s.db.GetDb())

	appGroupUsecase := appGroupUsecase.NewAppGroupUsecase(appGroupGormRepository)
	appGroupHandler := appGroupHandler.NewAppGroupHttpHandler(appGroupUsecase)

	// Routers
	appGroupRouter.SetAppGroupRoutes(api, appGroupHandler)
}
