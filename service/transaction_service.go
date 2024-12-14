package service

import (
	"errors"

	"github.com/firasam/julo/dto"
	"github.com/firasam/julo/model"
	r "github.com/firasam/julo/repository"
)

type TransactionService interface {
	ViewTransaction(xID string) ([]*model.Transaction, error)
	Deposit(xID string, input *dto.DepositRequest) (*model.Transaction, error)
	Withdraw(xID string, input *dto.WithdrawRequest) (*model.Transaction, error)
}

type transactionService struct {
	transactionRepository r.TransactionRepository
	walletRepository      r.WalletRepository
}

type TransactionConfig struct {
	TransactionRepository r.TransactionRepository
	WalletRepository      r.WalletRepository
}

func NewTransactionService(c *TransactionConfig) TransactionService {
	return &transactionService{
		transactionRepository: c.TransactionRepository,
		walletRepository:      c.WalletRepository,
	}
}

func (s *transactionService) ViewTransaction(xID string) ([]*model.Transaction, error) {
	wallet, err := s.walletRepository.FindByXID(xID)
	if !wallet.Enabled {
		return nil, errors.New("Transaction disabled")
	}
	transaction, err := s.transactionRepository.FindByXID(xID)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s *transactionService) Deposit(xID string, input *dto.DepositRequest) (*model.Transaction, error) {
	wallet, err := s.walletRepository.FindByXID(xID)
	if err != nil {
		return nil, err
	}
	if !wallet.Enabled {
		return nil, errors.New("Transaction disabled")
	}
	wallet.Balance += input.Amount

	transaction := model.Transaction{}
	transaction.Amount = input.Amount
	transaction.CustomerXid = xID
	transaction.TransactionType = "Deposit"
	transaction.Status = "Success"
	transaction.ReferenceID = input.ReferenceID
	_, err = s.transactionRepository.Save(&transaction)
	if err != nil {
		return nil, err
	}

	_, err = s.walletRepository.Update(wallet)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (s *transactionService) Withdraw(xID string, input *dto.WithdrawRequest) (*model.Transaction, error) {
	wallet, err := s.walletRepository.FindByXID(xID)
	if err != nil {
		return nil, err
	}
	if !wallet.Enabled {
		return nil, errors.New("Transaction disabled")
	}

	transaction := model.Transaction{}
	transaction.Amount = input.Amount
	transaction.CustomerXid = xID
	transaction.TransactionType = "Withdraw"
	transaction.Status = "Success"
	transaction.ReferenceID = input.ReferenceID
	_, err = s.transactionRepository.Save(&transaction)
	if err != nil {
		return nil, err
	}

	wallet.Balance -= input.Amount

	_, err = s.walletRepository.Update(wallet)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
