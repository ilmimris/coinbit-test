package registry

import (
	"github.com/ilmimris/coinbit-test/config"
	"github.com/ilmimris/coinbit-test/internal/core/port/inbound/service"
	coresvc "github.com/ilmimris/coinbit-test/internal/core/service"
)

func (r *registry) NewAccountService() service.WalletService {
	return coresvc.NewWalletService(
		coresvc.SetBalanceViewTable(*r.viewTable),
		coresvc.SetThresholdViewTable(*r.viewTable),
		coresvc.SetThreshold(int64(config.GlobalConfig.Wallet.Threshold)),
		coresvc.SetRollingPeriod(int(config.GlobalConfig.Wallet.RollingPeriod)),
		coresvc.SetDepositTopicPublisher(*r.producer),
	)
}
