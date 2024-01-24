package models

type ConfirmationCodeRequest struct {
}

type ConfirmationCodeResponse struct {
	Attempts int64  `json:"attempts"`
	Code     string `json:"code"`
}
