package route

import s "github.com/firasam/julo/service"

type Router struct {
	walletService      s.WalletService
	transactionService s.TransactionService
	jwtService         s.JWTService
}

type RouterConfig struct {
	WalletService      s.WalletService
	TransactionService s.TransactionService
	JWTService         s.JWTService
}

func NewRouter(c *RouterConfig) *Router {
	return &Router{
		walletService:      c.WalletService,
		transactionService: c.TransactionService,
		jwtService:         c.JWTService,
	}
}
