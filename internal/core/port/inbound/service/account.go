package service

type AccountService interface {
	Deposit(walletID string, amount float64) error
	Balance(walletID string) (float64, error)
}
