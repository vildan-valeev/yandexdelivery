package models

type ApplyChangesRequest struct {
	Changes           []Change `json:"changes"`
	LastKnownRevision string   `json:"last_known_revision"`
}

type ApplyChangesResponse struct {
}

type Change struct {
	Comment        string  `json:"comment"`
	Contact        Contact `json:"contact"`
	DropOffPointID int64   `json:"dropoff_point_id"`
	Items          []Item  `json:"items"`
	Kind           string  `json:"kind"`
	PickupPointID  int64   `json:"pickup_point_id"`
	PointID        int64   `json:"point_id"`
}
