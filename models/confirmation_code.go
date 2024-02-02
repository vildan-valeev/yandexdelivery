package models

type ConfirmationCodeRequest struct {
	ClaimID string `json:"claim_id"`
}

type ConfirmationCodeResponse struct {
	Attempts int64  `json:"attempts"`
	Code     string `json:"code"`
}
