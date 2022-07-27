package repository

type WalletRepository interface {
	SaveDeposit(walletID string, amount float64) error
	FetchBalance(walletID string) (float64, error)
}
