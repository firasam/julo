package dto

type WalletCreateRequest struct {
	CustomerXID string `form:"costumer_xid"`
}

type WalletResponse struct {
	Data Token `json:"data"`
	Status
}

type Token struct {
	Token string `json:"token"`
}

func FormatCreatedDetail(token string) WalletResponse {
	formattedWalletResponse := WalletResponse{}
	formattedWalletResponse.Data = Token{Token: token}
	formattedWalletResponse.Status.Status = "Success"
	return formattedWalletResponse
}
