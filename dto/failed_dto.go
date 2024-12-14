package dto

type Failed struct {
	Status string `json:"status"`
	Data   Error  `json:"data"`
}

type Error struct {
	Error string `json:"error"`
}

func FormatFail(err string) Failed {
	formatedFail := Failed{}
	formatedFail.Status = "failed"
	formatedFail.Data = Error{Error: err}
	return formatedFail
}
