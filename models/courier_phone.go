package models

type CourierPhoneRequest struct {
	ClaimID string `json:"claim_id"`
	PointID int64  `json:"point_id"`
}

type CourierPhoneResponse struct {
	Ext        string `json:"ext"`
	Phone      string `json:"phone"`
	TTLSeconds int64  `json:"ttl_seconds"`
}
