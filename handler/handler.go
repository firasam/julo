package handler

import s "github.com/firasam/julo/service"

type Handler struct {
	walletService      s.WalletService
	transactionService s.TransactionService
	jwtService         s.JWTService
}

type HandlerConfig struct {
	WalletService      s.WalletService
	TransactionService s.TransactionService
	JWTService         s.JWTService
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		walletService:      c.WalletService,
		transactionService: c.TransactionService,
		jwtService:         c.JWTService,
	}
}
