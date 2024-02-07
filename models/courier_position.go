package models

type CourierPositionResponse struct {
	Position    CourierPosition `json:"position"`
	RoutePoints []CourierPoint  `json:"route_points"`
}

type CourierPoint struct {
	ID          int64  `json:"id"`
	SharingLink string `json:"sharing_link"`
	Type        string `json:"type"`
	VisitOrder  int64  `json:"visit_order"`
}

type CourierPosition struct {
	Accuracy  float64 `json:"accuracy"`
	Number    float64 `json:"number"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Speed     float64 `json:"speed"`
	Timestamp int64   `json:"timestamp"`
}
