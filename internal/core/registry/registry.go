package registry

import (
	consumerRegistry "github.com/ilmimris/coinbit-test/internal/adapter/inbound/consumer/registry"
	restRegistry "github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/registry"
	"github.com/ilmimris/coinbit-test/internal/core/port/outbond/pubsub"
)

type registry struct {
	subscriber *pubsub.Subscriber
	producer   *pubsub.Publisher
	viewTable  *pubsub.ViewTable
}

type Registry interface {
	NewRestRegistryService() *restRegistry.ServiceRegistry
	NewConsumerRegistryService() *consumerRegistry.ServiceRegistry
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

func (r *registry) NewConsumerRegistryService() *consumerRegistry.ServiceRegistry {
	return consumerRegistry.NewServiceRegistry(
		consumerRegistry.NewAccountService(r.NewAccountService()),
	)
}

func NewGokaProducer(p *pubsub.Publisher) Option {
	return func(r *registry) {
		r.producer = p
	}
}

func NewGokaViewTable(v *pubsub.ViewTable) Option {
	return func(r *registry) {
		r.viewTable = v
	}
}

func NewGokaConsumerGroup(s *pubsub.Subscriber) Option {
	return func(r *registry) {
		r.subscriber = s
	}
}
