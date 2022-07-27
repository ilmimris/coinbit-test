package service

import (
	"context"

	"github.com/ilmimris/coinbit-test/internal/core/domain"
	"github.com/lovoo/goka"
)

type WalletService interface {
	Deposit(ctx context.Context, walletID string, amount float64) error
	AddBalance(ctx goka.Context, walletID string, amount float64) (float64, error)
	ProcessThreshold(ctx goka.Context, walletID string, amount float64) error
	GetDetail(ctx context.Context, walletID string) (domain.Wallet, error)
}
