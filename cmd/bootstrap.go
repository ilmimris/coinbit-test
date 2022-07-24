package cmd

import (
	"os"
	"os/signal"

	"github.com/ilmimris/coinbit-test/config"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/interfaces"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/registry"
)

type bootstrap struct {
	rest        interfaces.Rest
	optServices []OptionService
}

type OptionService func(b *bootstrap)

type Bootstrapper interface {
	Initialize(cfgFile string)
	AddServices(services ...OptionService)
	Close()
	GetRegistryRest() *registry.ServiceRegistry
	GetRest() interfaces.Rest
	RunServices()
}

func (b *bootstrap) Initialize(cfgFile string) {
	config.LoadConfig(cfgFile)
}

func (b *bootstrap) AddServices(services ...OptionService) {
	b.optServices = append(b.optServices, services...)
}

func (b *bootstrap) RunServices() {
	for _, value := range b.optServices {
		value(b)
	}

	// add graceful shutdown when interrupt signal detected
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		b.Close()
	}()
}

func (b *bootstrap) Close() {
	b.rest.Close()
}

func (b *bootstrap) GetRegistryRest() *registry.ServiceRegistry {
	return registry.NewServiceRegistry().NewRestRegistryService()
}

func (b *bootstrap) GetRest() interfaces.Rest {
	return b.rest
}
