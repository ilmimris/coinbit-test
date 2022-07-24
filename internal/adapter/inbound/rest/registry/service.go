package registry

import "github.com/ilmimris/coinbit-test/internal/core/port/inbound/service"

type ServiceRegistry struct {
	accountService service.AccountService
}

type OptServiceRegistry func(*ServiceRegistry)

func NewServiceRegistry(option ...OptServiceRegistry) *ServiceRegistry {
	r := &ServiceRegistry{}

	for _, o := range option {
		o(r)
	}

	return r
}

func NewAccountService(svc service.AccountService) OptServiceRegistry {
	return func(r *ServiceRegistry) {
		r.accountService = svc
	}
}

func (s *ServiceRegistry) GetAccountService() service.AccountService {
	return s.accountService
}
