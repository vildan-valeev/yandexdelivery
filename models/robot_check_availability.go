package models

type RobotCheckAvailabilityRequest struct {
	RoutePoints []RobotRoutePoint `json:"route_points"`
}

type RobotCheckAvailabilityResponse struct {
	Available bool `json:"available"`
}

type RobotRoutePoint struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}
