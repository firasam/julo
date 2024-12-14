package repository

import (
	"github.com/firasam/julo/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindByXID(xid string) ([]*model.Transaction, error)
	Save(Transaction *model.Transaction) (*model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

type TransactionConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(c *TransactionConfig) TransactionRepository {
	return &transactionRepository{
		db: c.DB,
	}
}

func (r *transactionRepository) FindByXID(xid string) ([]*model.Transaction, error) {
	var Transaction []*model.Transaction

	err := r.db.Where("customer_xid = ?", xid).Find(&Transaction).Error
	if err != nil {
		return Transaction, err
	}

	return Transaction, nil
}

func (r *transactionRepository) Save(transaction *model.Transaction) (*model.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
