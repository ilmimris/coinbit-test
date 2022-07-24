package rest

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/group"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/handler"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/registry"
)

type Options struct {
	Port string
}

type Rest struct {
	router  *fiber.App
	options *Options
}

func NewRest(o *Options, serviceRegistry *registry.ServiceRegistry) *Rest {
	app := fiber.New()

	app.Use(recover.New())
	// generate id for every request
	app.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	r := &Rest{
		router:  app,
		options: o,
	}

	v1 := handler.V1New(handler.OptHandlerV1{
		ServiceRegistry: serviceRegistry,
	})

	group.InitRouterV1(r, *v1)

	return r

}

func (r *Rest) Serve() {
	log.Printf("API Listening on %s", r.options.Port)
	if err := r.router.Listen(r.options.Port); err != nil {
		panic(err)
	}
}

func (r *Rest) GetRouter() *fiber.App {
	return r.router
}

func (r *Rest) Close() {
	log.Printf("server gracefully shutting down...")
	err := r.router.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}
