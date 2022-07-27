package consumer

import (
	"context"
	"log"

	"github.com/ilmimris/coinbit-test/internal/core/port/outbond/pubsub"
	"github.com/lovoo/goka"
)

type GokaConsumerGroup struct {
	processor *goka.Processor
}

type GokaConsumerGroupOpt struct {
	Brokers            []string
	Group              string
	Topic              string
	Handler            pubsub.GokaEventHandler
	TopicManagerConfig *goka.TopicManagerConfig
	InputCodec         pubsub.GokaCodec
	TableCodec         pubsub.GokaCodec
}

func NewGokaConsumerGroup(opt GokaConsumerGroupOpt) (subscriber pubsub.Subscriber, err error) {

	g := goka.DefineGroup(goka.Group(opt.Group), goka.Input(goka.Stream(opt.Topic), opt.InputCodec, opt.Handler.Handle), goka.Persist(opt.TableCodec))

	p, err := goka.NewProcessor(opt.Brokers, g, goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(opt.TopicManagerConfig)), goka.WithConsumerGroupBuilder(goka.DefaultConsumerGroupBuilder))
	if err != nil {
		return
	}

	if err != nil {
		return
	}
	subscriber = &GokaConsumerGroup{
		processor: p,
	}
	return
}

func (c *GokaConsumerGroup) Subscribe() {
	ctx := context.Background()
	go func() {
		if err := c.processor.Run(ctx); err != nil {
			log.Fatalf("error running processor: %v", err)
		}
	}()
}

func (c *GokaConsumerGroup) Close() (err error) {
	c.processor.Stop()
	log.Println("closing consumer group")
	return
}
