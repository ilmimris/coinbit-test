package bootstrap

import (
	"log"
	"os"
	"os/signal"

	"github.com/ilmimris/coinbit-test/config"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/consumer"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/interfaces"
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest"
	"github.com/ilmimris/coinbit-test/internal/adapter/outbound/producer"
	"github.com/ilmimris/coinbit-test/internal/adapter/outbound/viewtable"
	"github.com/lovoo/goka"

	consumerRegistry "github.com/ilmimris/coinbit-test/internal/adapter/inbound/consumer/registry"
	restRegistry "github.com/ilmimris/coinbit-test/internal/adapter/inbound/rest/registry"
	"github.com/ilmimris/coinbit-test/internal/core/port/outbond/pubsub"
	"github.com/ilmimris/coinbit-test/internal/core/registry"
)

type bootstrap struct {
	rest        interfaces.Rest
	optServices []OptionService
	subscriber  pubsub.Subscriber
	producer    pubsub.Publisher
	viewTable   pubsub.ViewTable
}

type OptionService func(b *bootstrap)

type Bootstrapper interface {
	Initialize(cfgURL string)
	AddServices(services ...OptionService)
	Close()
	GetTopicManagerConfig() *goka.TopicManagerConfig
	GetRegistryRest() *restRegistry.ServiceRegistry
	GetRegistryConsumer() *consumerRegistry.ServiceRegistry
	GetRest() interfaces.Rest
	RunServices()
}

func (b *bootstrap) Initialize(cfgURL string) {
	log.Println("Initializing configuration: %s", cfgURL)
	config.ReadModuleConfig(cfgURL)
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

func (b *bootstrap) GetTopicManagerConfig() *goka.TopicManagerConfig {
	tmc := goka.NewTopicManagerConfig()
	// if multiple broker we can configure above 1
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1

	return tmc
}

func (b *bootstrap) GetRegistryRest() *restRegistry.ServiceRegistry {
	return registry.NewRegistry(
		registry.NewGokaConsumerGroup(&b.subscriber),
		registry.NewGokaProducer(&b.producer),
		registry.NewGokaViewTable(&b.viewTable),
	).NewRestRegistryService()
}

func (b *bootstrap) GetRest() interfaces.Rest {
	return b.rest
}

func (b *bootstrap) GetRegistryConsumer() *consumerRegistry.ServiceRegistry {
	return registry.NewRegistry(
		registry.NewGokaConsumerGroup(&b.subscriber),
		registry.NewGokaProducer(&b.producer),
		registry.NewGokaViewTable(&b.viewTable),
	).NewConsumerRegistryService()
}

func NewGokaConsumerGroup(params consumer.GokaConsumerGroupOpt) OptionService {
	return func(b *bootstrap) {
		subscriber, err := consumer.NewGokaConsumerGroup(params)
		if err != nil {
			panic(err)
		}

		b.subscriber = subscriber
	}
}

func NewGokaProducer(params producer.GokaProducerOpt) OptionService {
	return func(b *bootstrap) {
		producer, err := producer.NewGokaProducer(params)
		if err != nil {
			panic(err)
		}

		b.producer = producer
	}
}

func NewGokaViewTable(params viewtable.GokaViewTableOpt) OptionService {
	return func(b *bootstrap) {
		viewTable, err := viewtable.NewGokaViewTable(params)
		if err != nil {
			panic(err)
		}

		b.viewTable = viewTable
	}
}

func NewServiceRest(param rest.Options) OptionService {
	return func(b *bootstrap) {
		b.rest = rest.NewRest(&param, b.GetRegistryRest())
	}
}

func New() Bootstrapper {
	bst := &bootstrap{}
	return bst
}
