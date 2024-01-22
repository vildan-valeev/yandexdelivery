package models

import (
	"errors"
)

type InfoResponse struct {
	AvailableCancelState string             `json:"available_cancel_state"`
	CallbackProperties   CallbackProperties `json:"callback_properties"`
	CarrierInfo          CarrierInfo        `json:"carrier_info"`
	ClientRequirements   ClientRequirements `json:"client_requirements"`
	Comment              string             `json:"comment"`
	CorpClientID         string             `json:"corp_client_id"`
	CreatedTs            string             `json:"created_ts"`
	CurrentPointID       int64              `json:"current_point_id"`
	Due                  string             `json:"due"`
	EmergencyContact     EmergencyContact   `json:"emergency_contact"`
	ErrorMessages        ErrorMessages      `json:"error_messages"`
	Eta                  int64              `json:"eta"`
	ID                   string             `json:"id"`
	Items                Items              `json:"items"`
	MatchedCars          MatchedCars        `json:"matched_cars"`
	OptionalReturn       bool               `json:"optional_return"`
	PerformerInfo        PerformerInfo      `json:"performer_info"`
	Pricing              Pricing            `json:"pricing"`
	Revision             int64              `json:"revision"`
	RouteID              string             `json:"route_id"`
	RoutePoints          RoutePointsInfo    `json:"route_points"`
	SameDayData          SameDayData        `json:"same_day_data"`
	ShippingDocument     string             `json:"shipping_document"`
	SkipAct              bool               `json:"skip_act"`
	SkipClientNotify     bool               `json:"skip_client_notify"`
	SkipDoorToDoor       bool               `json:"skip_door_to_door"`
	SkipEmergencyNotify  bool               `json:"skip_emergency_notify"`
	Status               Status             `json:"status"`
	TaxiOffer            TaxiOffer          `json:"taxi_offer"`
	UpdatedTs            string             `json:"updated_ts"`
	UserRequestRevision  string             `json:"user_request_revision"`
	Version              int64              `json:"version"`
	Warnings             Warnings           `json:"warnings"`
}

type CallbackProperties struct {
	CallbackURL string `json:"callback_url"`
}

type CarrierInfo struct {
	Address string `json:"address"`
	Inn     string `json:"inn"`
	Name    string `json:"name"`
}

type ClientRequirements struct {
	AssignRobot  bool     `json:"assign_robot,omitempty"`
	CargoLoaders int64    `json:"cargo_loaders,omitempty"`
	CargoOptions []string `json:"cargo_options,omitempty"`
	CargoType    string   `json:"cargo_type,omitempty"`
	ProCourier   bool     `json:"pro_courier,omitempty"`
	TaxiClass    string   `json:"taxi_class"`
}

type EmergencyContact struct {
	Name                string `json:"name"`
	Phone               string `json:"phone"`
	PhoneAdditionalCode string `json:"phone_additional_code,omitempty"`
}
type ErrorMessages []ErrorMessage

type ErrorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Items []Item

type Item struct {
	CostCurrency  Currency       `json:"cost_currency"`
	CostValue     string         `json:"cost_value"`
	DropoffPoint  int64          `json:"droppof_point"`
	ExtraID       string         `json:"extra_id"`
	Fiscalization *Fiscalization `json:"fiscalization"`
	PickupPoint   int64          `json:"pickup_point"`
	Quantity      int64          `json:"quantity"`
	Size          *Size          `json:"size,omitempty"`
	Title         string         `json:"title"`
	Weight        float64        `json:"weight"`
}

type Currency string

const (
	RUB Currency = "RUB"
)

func ToCurrency(s string) (Currency, error) {
	c := Currency(s)
	switch c {
	case RUB:
		return RUB, nil
	default:
		return "", errors.New("invalid Currency")
	}
}

type Fiscalization struct {
	Article     string `json:"article,omitempty"`
	Excise      string `json:"excise,omitempty"`
	ItemType    string `json:"item_type,omitempty"`
	Mark        *Mark  `json:"mark,omitempty"`
	SupplierInn string `json:"supplier_inn,omitempty"`
	VatCodeStr  string `json:"vat_code_str,omitempty"`
}

type Mark struct {
	Code string `json:"code"`
	Kind string `json:"kind"`
}

type Size struct {
	Height float64 `json:"height,omitempty"`
	Length float64 `json:"length,omitempty"`
	Width  float64 `json:"width,omitempty"`
}

type MatchedCars []MatchedCar

type MatchedCar struct {
	CargoLoaders    int64  `json:"cargo_loaders"`
	CargoType       string `json:"cargo_type"`
	CargoTypeInt    int64  `json:"cargo_type_int"`
	ClientTaxiClass string `json:"client_taxi_class"`
	DoorToDoor      bool   `json:"door_to_door"`
	ProCourier      bool   `json:"pro_courier"`
	TaxiClass       string `json:"taxi_class"`
}

type PerformerInfo struct {
	CarColor      string `json:"car_color"`
	CarColorHex   string `json:"car_color_hex"`
	CarModel      string `json:"car_model"`
	CarNumber     string `json:"car_number"`
	CourierName   string `json:"courier_name"`
	LegalName     string `json:"legal_name"`
	TransportType string `json:"transport_type"`
}

type Pricing struct {
	Currency           string        `json:"currency"`
	CurrencyRules      CurrencyRules `json:"currency_rules"`
	FinalPrice         string        `json:"final_price"`
	FinalPricingCalcID string        `json:"final_pricing_calc_id"`
	Offer              Offer         `json:"offer"`
}

type CurrencyRules struct {
	Code     string `json:"code"`
	Sign     string `json:"sign"`
	Template string `json:"template"`
	Text     string `json:"text"`
}

type Offer struct {
	OfferID      string `json:"offer_id"`
	Price        string `json:"price"`
	PriceRaw     int64  `json:"price_raw"`
	PriceWithVat string `json:"price_with_vat"`
	ValidUntil   string `json:"valid_until"`
}

type RoutePointsInfo []RoutePointInfo

type RoutePointInfo struct {
	Address               Address                `json:"address"`
	Contact               Contact                `json:"contact"`
	ExpectedVisitInterval *ExpectedVisitInterval `json:"expected_visit_interval"`
	ExternalOrderCost     *ExternalOrderCost     `json:"external_order_cost"`
	ExternalOrderID       string                 `json:"external_order_id,omitempty"`
	ID                    int64                  `json:"id,omitempty"`
	LeaveUnderDoor        bool                   `json:"leave_under_door,omitempty"`
	MeetOutside           bool                   `json:"meet_outside,omitempty"`
	ModifierAgeCheck      bool                   `json:"modifier_age_check,omitempty"`
	NoDoorCall            bool                   `json:"no_door_call,omitempty"`
	PaymentOnDelivery     *PaymentOnDelivery     `json:"payment_on_delivery,omitempty"`
	PerformerCancelReason string                 `json:"performer_cancel_reason,omitempty"`
	PickupCode            string                 `json:"pickup_code,omitempty"`
	ReturnComment         string                 `json:"return_comment,omitempty"`
	ReturnReasons         []string               `json:"return_reasons,omitempty"`
	SkipConfirmation      bool                   `json:"skip_confirmation,omitempty"`
	Type                  string                 `json:"type,omitempty"`
	VisitOrder            int64                  `json:"visit_order,omitempty"`
	VisitStatus           string                 `json:"visit_status,omitempty"`
	VisitedAt             VisitedAt              `json:"visited_at,omitempty"`
}

type Address struct {
	Building      string    `json:"building,omitempty"`
	BuildingName  string    `json:"building_name,omitempty"`
	City          string    `json:"city,omitempty"`
	Comment       string    `json:"comment,omitempty"`
	Coordinates   []float64 `json:"coordinates"`
	Country       string    `json:"country,omitempty"`
	Description   string    `json:"description,omitempty"`
	DoorCode      string    `json:"door_code,omitempty"`
	DoorCodeExtra string    `json:"door_code_extra,omitempty"`
	DoorbellName  string    `json:"doorbell_name,omitempty"`
	Flat          int64     `json:"flat,omitempty"`
	Floor         int64     `json:"floor,omitempty"`
	FullName      string    `json:"fullname"`
	Porch         string    `json:"porch,omitempty"`
	SFlat         string    `json:"sflat,omitempty"`
	SFloor        string    `json:"sfloor,omitempty"`
	Shortname     string    `json:"shortname,omitempty"`
	Street        string    `json:"street,omitempty"`
	URI           string    `json:"uri,omitempty"`
}

type Contact struct {
	Email               string `json:"email,omitempty"`
	Name                string `json:"name,omitempty"`
	Phone               string `json:"phone,omitempty"`
	PhoneAdditionalCode string `json:"phone_additional_code"`
}

type ExpectedVisitInterval struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type ExternalOrderCost struct {
	Currency     string `json:"currency"`
	CurrencySign string `json:"currency_sign"`
	Value        string `json:"value"`
}

type PaymentOnDelivery struct {
	ClientOrderID string        `json:"client_order_id"`
	Cost          string        `json:"cost"`
	Customer      Customer      `json:"customer"`
	InvoiceLink   string        `json:"invoice_link"`
	IsPaid        bool          `json:"is_paid"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	PaymentRefID  string        `json:"payment_ref_id"`
}

type PaymentMethod string

const (
	PaymentMethodCash PaymentMethod = "cash"
	PaymentMethodCard PaymentMethod = "card"
)

func ToPaymentMethod(s string) (PaymentMethod, error) {
	p := PaymentMethod(s)
	switch p {
	case PaymentMethodCash, PaymentMethodCard:
		return p, nil
	default:
		return "", errors.New("invalid payment method")
	}
}

type Customer struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Inn      string `json:"inn"`
	Phone    string `json:"phone"`
}

type VisitedAt struct {
	Actual                 string `json:"actual"`
	Expected               string `json:"expected"`
	ExpectedWaitingTimeSec int64  `json:"expected_waiting_time_sec"`
}

type SameDayData struct {
	DeliveryInterval DeliveryInterval `json:"delivery_interval"`
}

type DeliveryInterval struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type TaxiOffer struct {
	OfferID  string `json:"offer_id"`
	Price    string `json:"price"`
	PriceRaw int64  `json:"price_raw"`
}

type Warnings []Warning

type Warning struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Source  string `json:"source"`
}
