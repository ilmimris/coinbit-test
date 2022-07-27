package group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/handler"
)

func InitV1Group(root fiber.Router, h handler.HandlerV1) fiber.Router {
	v1Group := root.Group("/v1")

	walletGroup := v1Group.Group("/wallet")
	walletGroup.Get("/", h.GetBalance())
	walletGroup.Post("/", h.SaveDeposit())
	walletGroup.Get("/h", handler.Healthcheck())

	return v1Group
}
