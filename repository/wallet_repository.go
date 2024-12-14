package repository

import (
	"github.com/firasam/julo/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	FindByXID(xid string) (*model.Wallet, error)
	Save(wallet *model.Wallet) (*model.Wallet, error)
	Update(wallet *model.Wallet) (*model.Wallet, error)
}

type walletRepository struct {
	db *gorm.DB
}

type WalletConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(c *WalletConfig) WalletRepository {
	return &walletRepository{
		db: c.DB,
	}
}

func (r *walletRepository) FindByXID(xid string) (*model.Wallet, error) {
	var wallet *model.Wallet

	err := r.db.Where("customer_xid = ?", xid).First(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (r *walletRepository) Save(wallet *model.Wallet) (*model.Wallet, error) {
	err := r.db.Create(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (r *walletRepository) Update(wallet *model.Wallet) (*model.Wallet, error) {
	err := r.db.Save(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}
