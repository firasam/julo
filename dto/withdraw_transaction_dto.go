package dto

import (
	"github.com/firasam/julo/model"
)

type WithdrawRequest struct {
	Amount      int    `form:"amount"`
	ReferenceID string `form:"reference_id"`
}

type WithdrawTransactionResponse struct {
	Status
	Data Withdraw `json : "data"`
}

type Withdraw struct {
	Withdraw Transaction `json : "withdraw"`
}

func FormatWithdrawResponse(transactions *model.Transaction) WithdrawTransactionResponse {
	singleTransaction := Transaction{}
	singleTransaction.Id = transactions.ID
	singleTransaction.Status = transactions.Status
	singleTransaction.Transacted_at = transactions.CreatedAt
	singleTransaction.Type = transactions.TransactionType
	singleTransaction.Amount = transactions.Amount
	singleTransaction.ReferenceID = transactions.ReferenceID

	withdrawTransactionResponse := WithdrawTransactionResponse{}
	withdrawTransactionResponse.Status.Status = "Success"
	withdrawTransactionResponse.Data.Withdraw = singleTransaction
	return withdrawTransactionResponse
}
