package models

type TariffsRequest struct {
	FullName   string  `json:"fullname"`
	StartPoint []int64 `json:"start_point"`
}

type TariffsResponse struct {
	AvailableTariffs []Tariff `json:"available_tariffs"`
}
