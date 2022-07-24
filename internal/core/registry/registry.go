package registry

import (
	restRegistry "github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/registry"
	repoRegistry "github.com/ilmimris/coinbit-test/internal/core/registry/repository"
)

type registry struct{}

type Registry interface {
	NewRestRegistryService() *restRegistry.ServiceRegistry
}

type Option func(*registry)

func NewRegistry(option ...Option) Registry {
	r := &registry{}

	for _, o := range option {
		o(r)
	}

	return r
}

func (r *registry) NewRestRegistryService() *restRegistry.ServiceRegistry {
	return restRegistry.NewServiceRegistry(
		restRegistry.NewAccountService(r.NewAccountService()),
	)
}

func (r *registry) NewRepoRegistry() repoRegistry.RepositoryRegistry {
	return repositories.NewRepositoryRegistry(repositories.OptRepoRegistry{})
}
