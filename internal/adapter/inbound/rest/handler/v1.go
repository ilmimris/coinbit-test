package handler

import "github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/registry"

type HandlerV1 struct {
	serviceRegistry *registry.ServiceRegistry
}

type OptHandlerV1 struct {
	ServiceRegistry *registry.ServiceRegistry
}

func V1New(option OptHandlerV1) *HandlerV1 {
	return &HandlerV1{
		serviceRegistry: option.ServiceRegistry,
	}
}

func (h *HandlerV1) GetServiceRegistry() *registry.ServiceRegistry {
	return h.serviceRegistry
}
