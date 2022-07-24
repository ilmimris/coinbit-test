package service

import (
	"log"

	"github.com/ilmimris/coinbit-test/internal/core/port/inbound/service"
	repoRegistry "github.com/ilmimris/coinbit-test/internal/core/registry/repository"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

const (
	DepositTopic = "deposits"
)

type accountService struct {
	repositoryRegistry repoRegistry.RepositoryRegistry
	brokers            []string
}

type OptAccountSvc func(a *accountService)

func NewAccountService(options ...OptAccountSvc) service.AccountService {
	accountSvc := &accountService{}

	for _, option := range options {
		option(accountSvc)
	}

	return accountSvc
}

func NewAccountRepoRegistry(repoRegistry repoRegistry.RepositoryRegistry) OptAccountSvc {
	return func(a *accountService) {
		a.repositoryRegistry = repoRegistry
	}
}

func (a *accountService) Deposit(walletID string, amount float64) error {
	// goka emit Deposit event
	emitter, err := goka.NewEmitter(a.brokers, DepositTopic, new(codec.String))
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
		return err
	}

	defer emitter.Finish()

	err = emitter.EmitSync(walletID, amount)
	if err != nil {
		log.Fatalf("error emitting message: %v", err)
		return err
	}

	return nil
}

func (a *accountService) Balance(walletID string) (amount float64, err error) {
	return 0, nil
}
