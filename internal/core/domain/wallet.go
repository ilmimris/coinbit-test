package domain

type Wallet struct {
	WalletID string  `json:"wallet_id"`
	Amount   float64 `json:"amount"`
}
