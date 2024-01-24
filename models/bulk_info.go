package models

type BulkInfoRequest struct {
	ClaimIDs []string `json:"claim_ids"`
}

type BulkInfoResponse struct {
	Claims []Claim `json:"claims"`
}
