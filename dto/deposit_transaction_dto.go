package dto

import (
	"github.com/firasam/julo/model"
)

type DepositRequest struct {
	Amount      int    `form:"amount"`
	ReferenceID string `form:"refrence_id"`
}

type DepositTransactionResponse struct {
	Status
	Data Deposit `json : "data"`
}

type Deposit struct {
	Deposit Transaction `json : "deposit"`
}

func FormatDepositResponse(transactions *model.Transaction) DepositTransactionResponse {
	singleTransaction := Transaction{}
	singleTransaction.Id = transactions.ID
	singleTransaction.Status = transactions.Status
	singleTransaction.Transacted_at = transactions.CreatedAt
	singleTransaction.Type = transactions.TransactionType
	singleTransaction.Amount = transactions.Amount
	singleTransaction.ReferenceID = transactions.ReferenceID

	depositTransactionResponse := DepositTransactionResponse{}
	depositTransactionResponse.Status.Status = "Success"
	depositTransactionResponse.Data.Deposit = singleTransaction
	return depositTransactionResponse
}
