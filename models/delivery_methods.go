package models

type DeliveryMethodsRequest struct {
	FullName   string    `json:"full_name"`
	StartPoint []float64 `json:"start_point"`
}

type DeliveryMethodsResponse struct {
	ExpressDelivery ExpressDelivery `json:"express_delivery"`
	SameDayDelivery SameDayDelivery `json:"same_day_delivery"`
}

type ExpressDelivery struct {
	Allowed          bool     `json:"allowed"`
	AvailableTariffs []Tariff `json:"available_tariffs"`
}

type Tariff struct {
	MinimalPrice          int64               `json:"minimal_price"`
	Name                  string              `json:"name"`
	SupportedRequirements []TariffRequirement `json:"supported_requirements"`
	Text                  string              `json:"text"`
	Title                 string              `json:"title"`
}

type TariffRequirement struct {
	Default  bool                      `json:"default"`
	Name     string                    `json:"name"`
	Options  []TariffRequirementOption `json:"options"`
	Required bool                      `json:"required"`
	Text     string                    `json:"text"`
	Title    string                    `json:"title"`
	Type     string                    `json:"type"`
}

type TariffRequirementOption struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

type SameDayDelivery struct {
	Allowed            bool `json:"allowed"`
	AvailableIntervals []AvailableInterval
}

type AvailableInterval struct {
	From string `json:"from"`
	To   string `json:"to"`
}
