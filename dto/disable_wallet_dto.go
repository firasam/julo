package dto

import (
	"time"

	"github.com/firasam/julo/model"
)

type DisableResponse struct {
	Status
	Data DisableWallet `json : "data"`
}

type DisableWallet struct {
	Id          string    `json: "id"`
	Owned       string    `json: "owned_by"`
	Enabled     string    `json: "status"`
	Disabled_at time.Time `json: "disabled_at"`
	Balance     int       `json: "balance"`
}

func FormatedDisable(wallet *model.Wallet) DisableResponse {
	w := DisableWallet{}
	w.Id = wallet.ID
	w.Owned = wallet.CustomerXid
	w.Enabled = "disabled"
	w.Disabled_at = wallet.ChangedAt
	w.Balance = wallet.Balance
	response := DisableResponse{}
	response.Status.Status = "Success"
	response.Data = w
	return response
}
