package registry

import "github.com/ilmimris/coinbit-test/internal/core/port/inbound/service"

type ServiceRegistry struct {
	accountService service.WalletService
}

type OptServiceRegistry func(*ServiceRegistry)

func NewServiceRegistry(option ...OptServiceRegistry) *ServiceRegistry {
	r := &ServiceRegistry{}

	for _, o := range option {
		o(r)
	}

	return r
}

func NewAccountService(svc service.WalletService) OptServiceRegistry {
	return func(r *ServiceRegistry) {
		r.accountService = svc
	}
}

func (s *ServiceRegistry) GetAccountService() service.WalletService {
	return s.accountService
}
