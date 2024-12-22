package server

import (
	"fmt"
	"shopping-service-be/pkg/config"
	"shopping-service-be/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type (
	Server interface {
		Start()
	}

	fiberServer struct {
		app  *fiber.App
		db   database.Database
		conf *config.Config
	}
)

func NewFiberServer(conf *config.Config, db database.Database) Server {
	fiberApp := fiber.New(fiber.Config{
		ReadBufferSize:        60 * 1024,
		DisableStartupMessage: false,
	})

	return &fiberServer{
		app:  fiberApp,
		db:   db,
		conf: conf,
	}
}

func (s *fiberServer) Start() {
	s.app.Use(logger.New())

	api := s.app.Group("/api")
	s.app.Get("", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Good health ✅",
		})
	})

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.initializeProductHttpHandler(api)
	s.initializeUserHttpHandler(api)
	s.initializeAuthHttpHandler(api, s.conf)
	s.app.Listen(serverUrl)
}
