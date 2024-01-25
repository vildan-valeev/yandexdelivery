package models

type PhotosPointRequest struct {
	ClaimID      string `json:"claim_id"`
	ClaimPointID int64  `json:"claim_point_id"`
}

type PhotosPointResponse struct {
	Photos PhotoPoint `json:"photos"`
}

type PhotoPoint struct {
	ClaimPointID int64  `json:"claim_point_id"`
	ExternalID   string `json:"external_id"`
	ID           string `json:"id"`
	Status       Status `json:"status"`
	URL          string `json:"url"`
}
