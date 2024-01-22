package models

type DeliveryMethodsRequest struct {
	FullName   string    `json:"full_name"`
	StartPoint []float64 `json:"start_point"`
}
