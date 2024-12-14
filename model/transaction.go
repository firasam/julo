package model

type Transaction struct {
	Base
	CustomerXid     string
	TransactionType string `gorm:"unique"`
	Amount          int
	ReferenceID     string `gorm:"unique"`
	Status          string
}

func (Transaction) TableName() string {
	return "transactions"
}
