package models

type OffersCalculateRequest struct {
	Items        []DeliveryItem             `json:"items"`
	Requirements OffersCalculateRequirement `json:"requirements"`
	RoutePoints  []RoutePoint               `json:"route_points"`
}

type OffersCalculateResponse struct {
	Offers []CalculatedOffer `json:"offers"`
}

type CalculatedOffer struct {
	DeliveryInterval DeliveryInterval `json:"delivery_interval"`
	Description      string           `json:"description"`
	OfferTTL         string           `json:"offer_ttl"`
	Payload          string           `json:"payload"`
	PickupInterval   DeliveryInterval `json:"pickup_interval"`
	Price            OfferPrice       `json:"price"`
	TaxiClass        string           `json:"taxi_class"`
}

type OffersCalculateRequirement struct {
	CargoLoaders   int64    `json:"cargo_loaders"`
	CargoOptions   []string `json:"cargo_options"`
	CargoType      string   `json:"cargo_type"`
	Due            string   `json:"due"`
	ProCourier     bool     `json:"pro_courier"`
	SkipDoorToDoor bool     `json:"skip_door_to_door"`
	TaxiClasses    []string `json:"taxi_class"`
}

type OfferPrice struct {
	BasePrice         string `json:"base_price"`
	Currency          string `json:"currency"`
	SureRatio         int64  `json:"sure_ratio"`
	TotalPrice        string `json:"total_price"`
	TotalPriceWithVat string `json:"total_price_with_vat"`
}
