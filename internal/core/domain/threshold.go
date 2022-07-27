package domain

import "time"

type Threshold struct {
	WalletId              string    `json:"wallet_id"`
	Amount                float64   `json:"amount"`
	TotalAmountWithWindow float64   `json:"total_amount_with_window"`
	StartWindowTime       time.Time `json:"start_window_time"`
	CreatedTime           time.Time `json:"created_time"`
	AboveThreshold        bool      `json:"above_threshold"`
}
