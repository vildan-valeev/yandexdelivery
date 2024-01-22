package models

type DriverPhoneRequest struct {
	ClaimID string `json:"claim_id"`
	PointID int64  `json:"point_id"`
}

type DriverPhoneResponse struct {
	Ext        string `json:"ext"`
	Phone      string `json:"phone"`
	TTLSeconds int64  `json:"ttl_seconds"`
}
