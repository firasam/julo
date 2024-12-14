package dto

import (
	"time"

	"github.com/firasam/julo/model"
)

type ViewTransactionResponse struct {
	Status
	Data AllTransaction `json : "data"`
}

type AllTransaction struct {
	Transaction []Transaction `json : "data"`
}

type Transaction struct {
	Id            string    `json: "id"`
	Status        string    `json:"status"`
	Transacted_at time.Time `json:"transacted_at"`
	Type          string    `json:"Type"`
	Amount        int       `json:"Amount"`
	ReferenceID   string    `json:"reference_id"`
}

func FormatViewTransactionResponse(transactions []*model.Transaction) ViewTransactionResponse {
	responseTransaction := []Transaction{}
	for _, t := range transactions {
		singleTransaction := Transaction{}
		singleTransaction.Id = t.ID
		singleTransaction.Status = t.Status
		singleTransaction.Transacted_at = t.CreatedAt
		singleTransaction.Type = t.TransactionType
		singleTransaction.Amount = t.Amount
		singleTransaction.ReferenceID = t.ReferenceID
		responseTransaction = append(responseTransaction, singleTransaction)
	}
	viewTransactionResponse := ViewTransactionResponse{}
	viewTransactionResponse.Status.Status = "Success"
	viewTransactionResponse.Data.Transaction = responseTransaction
	return viewTransactionResponse
}
