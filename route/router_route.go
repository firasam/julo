package route

import (
	"github.com/firasam/julo/handler"
	"github.com/firasam/julo/middleware"
	"github.com/gin-gonic/gin"
)

func (r *Router) Wallet(route *gin.RouterGroup, h *handler.Handler) {
	route.POST("/init", h.Create)
	route.POST("/wallet", middleware.AuthMiddleware(r.jwtService), h.EnableWallet)
	route.PATCH("/wallet", middleware.AuthMiddleware(r.jwtService), h.DisableWallet)
	route.GET("/wallet", middleware.AuthMiddleware(r.jwtService), h.ViewWallet)

	route.GET("/wallet/transactions", middleware.AuthMiddleware(r.jwtService), h.ViewTransaction)
	route.POST("/wallet/deposits", middleware.AuthMiddleware(r.jwtService), h.Deposit)
	route.POST("/wallet/withdraws", middleware.AuthMiddleware(r.jwtService), h.Withdraw)
}
