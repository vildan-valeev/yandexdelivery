package models

type EditRequest struct {
	AutoAccept          string             `json:"auto_accept"`
	CallbackProperties  CallbackProperties `json:"callback_properties"`
	ClientRequirements  ClientRequirements `json:"client_requirements"`
	Comment             string             `json:"comment"`
	Due                 string             `json:"due"`
	EmergencyContact    EmergencyContact   `json:"emergency_contact"`
	Items               Items              `json:"items"`
	OfferPayload        string             `json:"offer_payload"`
	OptionalReturn      bool               `json:"optional_return"`
	ReferralSource      string             `json:"referral_source"`
	RoutePoints         []EditRoutePoint   `json:"route_points"`
	SameDayData         SameDayData        `json:"same_day_data"`
	ShippingDocument    string             `json:"shipping_document"`
	SkipAct             bool               `json:"skip_act"`
	SkipClientNotify    bool               `json:"skip_client_notify"`
	SkipDoorToDoor      bool               `json:"skip_door_to_door"`
	SkipEmergencyNotify bool               `json:"skip_emergency_notify"`
}

type EditRoutePoint struct {
	Address Address `json:"address"`
	Buyout  Buyout  `json:"buyout"`
}
