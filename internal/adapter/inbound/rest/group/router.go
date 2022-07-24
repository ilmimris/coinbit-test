package group

import (
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/interfaces"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/handler"
)

func InitRouterV1(rest interfaces.Rest, h handler.HandlerV1) {
	root := rest.GetRouter()
	root.Group("/")

	// init healthcheck
	InitHealthCheck(root)

	// define V1
	_ = InitV1Group(root, h)

}
