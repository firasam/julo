package main

import (
	"fmt"
	"os"

	"github.com/firasam/julo/db"
	"github.com/firasam/julo/handler"
	"github.com/firasam/julo/repository"
	"github.com/firasam/julo/route"
	"github.com/firasam/julo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.GetConn()

	walletRepository := repository.NewWalletRepository(&repository.WalletConfig{DB: db})
	transactionRepository := repository.NewTransactionRepository(&repository.TransactionConfig{DB: db})

	walletService := service.NewwalletService(&service.WalletConfig{WalletRepository: walletRepository})
	transactionService := service.NewTransactionService(&service.TransactionConfig{TransactionRepository: transactionRepository, WalletRepository: walletRepository})
	jwtService := service.NewJWTService(&service.JWTSConfig{})

	h := handler.NewHandler(&handler.HandlerConfig{
		WalletService:      walletService,
		TransactionService: transactionService,
		JWTService:         jwtService,
	})

	routes := route.NewRouter(&route.RouterConfig{WalletService: walletService, TransactionService: transactionService, JWTService: jwtService})

	router := gin.Default()

	version := os.Getenv("API_VERSION")
	api := router.Group(fmt.Sprintf("/api/%s", version))

	routes.Wallet(api, h)

	router.Run()
}
