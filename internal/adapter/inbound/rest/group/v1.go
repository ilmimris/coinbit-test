package group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/handler"
)

func InitV1Group(root fiber.Router, h handler.HandlerV1) fiber.Router {
	v1Group := root.Group("/v1")

	return v1Group
}
