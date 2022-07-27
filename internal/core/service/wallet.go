package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ilmimris/coinbit-test/internal/core/domain"
	"github.com/ilmimris/coinbit-test/internal/core/port/inbound/service"
	"github.com/ilmimris/coinbit-test/internal/core/port/outbond/pubsub"
	"github.com/lovoo/goka"
)

const (
	DepositTopic = "deposits"
)

type walletService struct {
	depositTopicPublisher pubsub.Publisher
	rollingPeriod         int
	threshold             int64
	balanceViewTable      pubsub.ViewTable
	thresholdViewTable    pubsub.ViewTable
}

type OptWalletSvc func(w *walletService)

func NewWalletService(options ...OptWalletSvc) service.WalletService {
	walletSvc := &walletService{}

	for _, option := range options {
		option(walletSvc)
	}

	return walletSvc
}

func SetRollingPeriod(p int) OptWalletSvc {
	return func(w *walletService) {
		w.rollingPeriod = p
	}
}

func SetThreshold(t int64) OptWalletSvc {
	return func(w *walletService) {
		w.threshold = t
	}
}

func SetDepositTopicPublisher(p pubsub.Publisher) OptWalletSvc {
	return func(w *walletService) {
		w.depositTopicPublisher = p
	}
}

func SetBalanceViewTable(vt pubsub.ViewTable) OptWalletSvc {
	return func(w *walletService) {
		w.balanceViewTable = vt
	}
}

func SetThresholdViewTable(vt pubsub.ViewTable) OptWalletSvc {
	return func(w *walletService) {
		w.thresholdViewTable = vt
	}
}

func (w *walletService) Deposit(ctx context.Context, walletID string, amount float64) error {
	var deposit = &domain.Wallet{
		WalletID: walletID,
		Amount:   amount,
	}

	// goka emit Deposit event
	err := w.depositTopicPublisher.Send(ctx, walletID, deposit)
	if err != nil {
		log.Fatalf("error publish deposit: %v", err)
		return err
	}

	return nil
}

func (w *walletService) AddBalance(ctx goka.Context, walletID string, amount float64) (float64, error) {
	var wallet *domain.Wallet
	if val := ctx.Value(); val != nil {
		wallet = val.(*domain.Wallet)
	} else {
		wallet = new(domain.Wallet)
	}

	wallet.WalletID = walletID
	wallet.Amount += amount
	ctx.SetValue(wallet)

	return wallet.Amount, nil
}

func (w *walletService) ProcessThreshold(ctx goka.Context, walletID string, amount float64) error {
	var threshold *domain.Threshold

	if val := ctx.Value(); val != nil {
		threshold = val.(*domain.Threshold)
	} else {
		threshold = new(domain.Threshold)
		threshold.StartWindowTime = time.Now()
	}

	now := time.Now()
	threshold.WalletId = walletID
	threshold.Amount = amount
	threshold.TotalAmountWithWindow += amount
	threshold.CreatedTime = now

	// get diff time between time when rolling period started and current deposit time
	diff := now.Sub(threshold.StartWindowTime).Seconds()

	if diff > float64(w.rollingPeriod) {
		// reset rolling period
		threshold.AboveThreshold = false
		threshold.StartWindowTime = now
		threshold.TotalAmountWithWindow = amount
	} else {
		// check if total amount with rolling period is above threshold
		if threshold.TotalAmountWithWindow > float64(w.threshold) {
			threshold.AboveThreshold = true
		} else {
			threshold.AboveThreshold = false
		}
	}
	ctx.SetValue(threshold)

	return nil
}

func (w *walletService) GetDetail(ctx context.Context, walletID string) (domain.Wallet, error) {
	balanceData, err := w.balanceViewTable.Get(walletID)
	if err != nil {
		log.Fatalf("error get balance data: %v", err)
		return domain.Wallet{}, err
	}

	thresholdData, err := w.thresholdViewTable.Get(walletID)
	if err != nil {
		log.Fatalf("error get threshold data: %v", err)
		return domain.Wallet{}, err
	}

	if balanceData == nil || thresholdData == nil {
		err := fmt.Errorf("wallet id %s not found", walletID)
		return domain.Wallet{}, err
	}

	balance := balanceData.(*domain.Wallet)
	threshold := thresholdData.(*domain.Threshold)

	detail := domain.Wallet{
		WalletID:       balance.WalletID,
		Amount:         balance.Amount,
		AboveThreshold: threshold.AboveThreshold,
	}

	return detail, nil
}
