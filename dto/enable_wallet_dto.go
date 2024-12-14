package dto

import (
	"time"

	"github.com/firasam/julo/model"
)

type EnabledResponse struct {
	Status
	Data Wallet `json : "data"`
}

type Wallet struct {
	Id         string    `json: "id"`
	Owned      string    `json: "owned_by"`
	Enabled    string    `json: "status"`
	Enabled_at time.Time `json: "enabled_at"`
	Balance    int       `json: "balance"`
}

func FormatedEnable(wallet *model.Wallet) EnabledResponse {
	w := Wallet{}
	w.Id = wallet.ID
	w.Owned = wallet.CustomerXid
	w.Enabled = "enabled"
	w.Enabled_at = wallet.ChangedAt
	w.Balance = wallet.Balance
	response := EnabledResponse{}
	response.Status.Status = "Success"
	response.Data = w
	return response
}
