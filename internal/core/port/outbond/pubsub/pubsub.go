package pubsub

import (
	"context"

	"github.com/lovoo/goka"
)

type EventHandler interface {
	Handle(ctx context.Context, message interface{}) error
}

type GokaEventHandler interface {
	Handle(ctx goka.Context, message interface{})
}

type Publisher interface {
	Send(ctx context.Context, key string, message interface{}) error
	Close() error
}

type Subscriber interface {
	Subscribe()
	Close() error
}

type ViewTable interface {
	Open()
	Get(key string) (data interface{}, err error)
	Close()
}

type GokaCodec interface {
	Encode(value interface{}) ([]byte, error)
	Decode(data []byte) (interface{}, error)
}

type GokaContext interface {
	goka.Context
}
