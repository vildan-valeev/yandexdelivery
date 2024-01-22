package models

type CheckPriceRequest struct {
	Items          []DeliveryItem        `json:"items"`
	Requirements   CheckPriceRequirement `json:"requirements"`
	RoutePoints    []RoutePoint          `json:"route_points"`
	SkipDoorToDoor bool                  `json:"skip_door_to_door"`
}

type DeliveryItem struct {
	DropoffPoint int64   `json:"droppof_point"`
	PickupPoint  int64   `json:"pickup_point"`
	Quantity     int64   `json:"quantity"`
	Size         *Size   `json:"size,omitempty"`
	Title        string  `json:"title"`
	Weight       float64 `json:"weight"`
}

type RoutePoint struct {
	Coordinates []float64 `json:"coordinates"`
	FullName    string    `json:"fullname"`
	ID          int64     `json:"id"`
}

type CheckPriceResponse struct {
	CurrencyRules  CurrencyRules         `json:"currency_rules"`
	DistanceMeters int64                 `json:"distance_meters"`
	ETA            int64                 `json:"eta"`
	Price          string                `json:"price"`
	Requirements   CheckPriceRequirement `json:"requirements"`
	ZoneID         string                `json:"zone_id"`
}

type CheckPriceRequirement struct {
	CargoLoaders int64       `json:"cargo_loaders"`
	CargoOptions []string    `json:"cargo_options"`
	CargoType    string      `json:"cargo_type"`
	ProCourier   bool        `json:"pro_courier"`
	SameDayData  SameDayData `json:"same_day_data"`
	TaxiClass    string      `json:"taxi_class"`
}
