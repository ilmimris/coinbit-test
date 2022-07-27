package registry

import "github.com/ilmimris/coinbit-test/internal/core/port/outbond/repository"

type RepositoryRegistry interface {
	GetAccountRepository() repository.WalletRepository
}
