package registry

import (
	"github.com/ilmimris/coinbit-test/internal/core/port/inbound/service"
	coresvc "github.com/ilmimris/coinbit-test/internal/core/service"
)

func (r *registry) NewAccountService() service.WalletService {
	return coresvc.NewWalletService()
}
