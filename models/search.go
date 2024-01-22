package models

type SearchRequest struct {
	ClaimID         string `json:"claim_id"`
	CreatedFrom     string `json:"created_from"`
	CreatedTo       string `json:"created_to"`
	Cursor          string `json:"cursor"`
	DueFrom         string `json:"due_from"`
	ExternalOrderID string `json:"external_order_id"`
	Limit           int64  `json:"limit"`
	Offset          int64  `json:"offset"`
	Phone           string `json:"phone"`
	State           string `json:"state"`
	Status          string `json:"status"`
}

type SearchResponse struct {
	Cursor string  `json:"cursor"`
	Claims []Claim `json:"claims"`
}
