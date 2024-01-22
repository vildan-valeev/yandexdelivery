package models

type TrackingLinksRequest struct {
}

type TrackingLinksResponse struct {
	RoutePoints []CourierPoint `json:"route_points"`
}
