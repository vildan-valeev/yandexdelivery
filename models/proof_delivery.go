package models

type ProofDeliveryResponse struct {
	ProofDelivery ProofDelivery `json:"proof_delivery"`
}

type ProofDelivery struct {
	ClaimPointID  int64         `json:"claim_point_id"`
	Photos        []PhotoPoint  `json:"photos"`
	RecipientInfo RecipientInfo `json:"recipient_info"`
}

type RecipientInfo struct {
	DocumentID    string `json:"document_id"`
	RecipientName string `json:"recipient_name"`
	RecipientType string `json:"recipient_type"`
}
