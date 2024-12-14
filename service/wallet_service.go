package service

import (
	"errors"
	"time"

	"github.com/firasam/julo/dto"
	"github.com/firasam/julo/model"
	r "github.com/firasam/julo/repository"
	"gorm.io/gorm"
)

type WalletService interface {
	CreateWallet(input *dto.WalletCreateRequest) (*model.Wallet, error)
	EnableWallet(xID string) (*model.Wallet, error)
	DisableWallet(xID string) (*model.Wallet, error)
	ViewWallet(xID string) (*model.Wallet, error)
}

type walletService struct {
	walletRepository r.WalletRepository
}

type WalletConfig struct {
	WalletRepository r.WalletRepository
}

func NewwalletService(c *WalletConfig) WalletService {
	return &walletService{
		walletRepository: c.WalletRepository,
	}
}

func (s *walletService) CreateWallet(input *dto.WalletCreateRequest) (*model.Wallet, error) {
	wallet, err := s.walletRepository.FindByXID(input.CustomerXID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return wallet, err
		} else {
			newWallet := model.Wallet{}
			newWallet.CustomerXid = input.CustomerXID
			_, err := s.walletRepository.Save(&newWallet)
			if err != nil {
				return &newWallet, err
			}

			return &newWallet, nil
		}
	}
	return wallet, nil
}

func (s *walletService) EnableWallet(xID string) (*model.Wallet, error) {
	wallet, err := s.walletRepository.FindByXID(xID)
	if err != nil {
		return nil, err
	}
	if wallet.Enabled {
		return nil, errors.New("already enabled")
	}
	wallet.Enabled = true
	wallet.ChangedAt = time.Now()
	s.walletRepository.Update(wallet)
	return wallet, nil
}

func (s *walletService) DisableWallet(xID string) (*model.Wallet, error) {
	wallet, err := s.walletRepository.FindByXID(xID)
	if err != nil {
		return nil, err
	}
	if !wallet.Enabled {
		return nil, errors.New("already disable")
	}
	wallet.Enabled = false
	wallet.ChangedAt = time.Now()
	s.walletRepository.Update(wallet)
	return wallet, nil
}

func (s *walletService) ViewWallet(xID string) (*model.Wallet, error) {
	wallet, err := s.walletRepository.FindByXID(xID)
	if err != nil {
		return nil, err
	}
	if !wallet.Enabled {
		return nil, errors.New("wallet disabled")
	}
	return wallet, nil
}
