package models

type JournalRequest struct {
	Cursor string `json:"cursor"`
}

type JournalResponse struct {
	Cursor string  `json:"cursor"`
	Events []Event `json:"events"`
}

type Event struct {
	ChangeType     string `json:"change_type"` // Возможные значения status_changed — изменение статуса; price_changed — изменение цены
	ClaimID        string `json:"claim_id"`
	ClaimOrigin    string `json:"claim_origin"`
	ClientID       string `json:"client_id"`
	CurrentPointID int64  `json:"current_point_id"`
	NewCurrency    string `json:"new_currency"`
	NewPrice       string `json:"new_price"`
	NewStatus      Status `json:"new_status"`
	OperationID    int64  `json:"operation_id"`
	Resolution     string `json:"resolution"`
	Revision       int64  `json:"revision"`
	UpdatedTS      string `json:"updated_ts"` // Время события в формате ISO 8601
}
