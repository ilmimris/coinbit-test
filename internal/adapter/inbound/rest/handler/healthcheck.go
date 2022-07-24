package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Healthcheck() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(200)
	}
}
