package group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/handler"
)

func InitHealthCheck(root fiber.Router) {
	healthCheckGroup := root.Group("/healthcheck")
	healthCheckGroup.Get("/", handler.Healthcheck())
}
