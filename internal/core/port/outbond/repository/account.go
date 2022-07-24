package repository

type AccountRepository interface {
	SaveDeposit(walletID string, amount float64) error
	FetchBalance(walletID string) (float64, error)
}
