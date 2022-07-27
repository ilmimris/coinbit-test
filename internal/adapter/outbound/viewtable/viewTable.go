package viewtable

import (
	"context"
	"log"

	"github.com/ilmimris/coinbit-test/internal/core/port/outbond/pubsub"
	"github.com/lovoo/goka"
)

type GokaViewTable struct {
	view   *goka.View
	cancel context.CancelFunc
}

type GokaViewTableOpt struct {
	Group   string
	Brokers []string
	Codec   pubsub.GokaCodec
}

func NewGokaViewTable(opt GokaViewTableOpt) (view pubsub.ViewTable, err error) {

	v, err := goka.NewView(opt.Brokers, goka.GroupTable(goka.Group(opt.Group)), opt.Codec)
	if err != nil {
		return nil, err
	}

	view = &GokaViewTable{
		view: v,
	}

	return
}

func (v *GokaViewTable) Open() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		if v.view.CurrentState() != goka.ViewStateRunning {
			if err := v.view.Run(ctx); err != nil {
				log.Fatalf("error running view: %v", err)
			}
		}
	}(ctx)
	v.cancel = cancel
	v.view.Run(context.Background())
}

func (v *GokaViewTable) Get(key string) (data interface{}, err error) {
	return v.view.Get(key)
}

func (v *GokaViewTable) Close() {
	v.cancel()
	log.Println("closing view table")
	return
}
