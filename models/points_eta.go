package models

//type PointsEtaRequest struct {
//}

type PointsEtaResponse struct {
	ID                string    `json:"id"`
	PerformerPosition []float64 `json:"performer_position"`
	RoutePoints       []RoutePointEta
}

type RoutePointEta struct {
	Address     Address   `json:"address"`
	ID          int64     `json:"id"`
	Type        string    `json:"type"`
	VisitOrder  int64     `json:"visit_order"`
	VisitStatus string    `json:"visit_status"`
	VisitedAt   VisitedAt `json:"visited_at"`
}
