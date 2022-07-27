package producer

import (
	"context"
	"log"

	"github.com/ilmimris/coinbit-test/internal/core/port/outbond/pubsub"
	"github.com/lovoo/goka"
)

type GokaProducer struct {
	emitter *goka.Emitter
}

type GokaProducerOpt struct {
	Brokers []string
	Topic   string
	Codec   pubsub.GokaCodec
}

func NewGokaProducer(opt GokaProducerOpt) (producer pubsub.Publisher, err error) {
	emitter, err := goka.NewEmitter(opt.Brokers, goka.Stream(opt.Topic), opt.Codec)
	if err != nil {
		return
	}
	producer = GokaProducer{
		emitter: emitter,
	}
	return
}

func (p GokaProducer) Send(ctx context.Context, key string, message interface{}) (err error) {
	return p.emitter.EmitSync(key, message)
}

func (p GokaProducer) Close() (err error) {
	err = p.emitter.Finish()
	if err == nil {
		log.Println("closing producer")
	}
	return
}
