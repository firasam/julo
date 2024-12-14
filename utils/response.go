package utils

import "time"

type WalletResponse struct {
	ID        string     `json:"id"`
	OwnedBy   string     `json:"owned_by"`
	Status    string     `json:"status"`
	EnabledAt *time.Time `json:"enabled_at"`
	Balance   int        `json:"balance"`
}
