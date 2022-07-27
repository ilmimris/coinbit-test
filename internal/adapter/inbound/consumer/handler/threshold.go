package handler

import (
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/consumer/registry"
	"github.com/ilmimris/coinbit-test/internal/core/domain"
	"github.com/ilmimris/coinbit-test/internal/core/port/outbond/pubsub"
	"github.com/lovoo/goka"
	log "github.com/sirupsen/logrus"
)

type ProcessThresholdEventHandler struct {
	serviceRegistry *registry.ServiceRegistry
}

func NewProcessThresholdEventHandler(service *registry.ServiceRegistry) pubsub.GokaEventHandler {
	return &ProcessThresholdEventHandler{
		serviceRegistry: service,
	}
}

func (h ProcessThresholdEventHandler) Handle(ctx goka.Context, message interface{}) {
	payload, ok := message.(*domain.Wallet)
	if !ok {
		log.Error("not a kafka message")
		return
	}

	err := h.serviceRegistry.GetAccountService().ProcessThreshold(ctx, payload.WalletID, payload.Amount)
	if err != nil {
		log.Info(err)
	}

}
