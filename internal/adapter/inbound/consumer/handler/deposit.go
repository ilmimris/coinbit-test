package handler

import (
	"github.com/ilmimris/coinbit-test/internal/adapter/inbound/consumer/registry"
	"github.com/ilmimris/coinbit-test/internal/core/domain"
	"github.com/ilmimris/coinbit-test/internal/core/port/outbond/pubsub"
	"github.com/lovoo/goka"
	log "github.com/sirupsen/logrus"
)

type DepositWalletEventHandler struct {
	serviceRegistry *registry.ServiceRegistry
}

func NewDepositWalletEventHandler(service *registry.ServiceRegistry) pubsub.GokaEventHandler {
	return &DepositWalletEventHandler{
		serviceRegistry: service,
	}
}

func (h DepositWalletEventHandler) Handle(ctx goka.Context, message interface{}) {
	payload, ok := message.(*domain.Wallet)
	if !ok {
		log.Error("not a kafka message")
		return
	}

	_, err := h.serviceRegistry.GetAccountService().AddBalance(ctx, payload.WalletID, payload.Amount)
	if err != nil {
		log.Info(err)
	}

}
