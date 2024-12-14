package model

import "time"

type Wallet struct {
	Base
	CustomerXid string
	Enabled     bool
	ChangedAt   time.Time
	Balance     int
}

func (Wallet) TableName() string {
	return "wallets"
}
