package models

import "devgit.apteka-aprel.ru/web/transporters/pkg/yandex_logistics/models/create"

type CreateRequest struct {
	AutoAccept          bool                `json:"auto_accept,omitempty"`
	CallbackProperties  *CallbackProperties `json:"callback_properties,omitempty"`
	ClientRequirements  ClientRequirements  `json:"client_requirements,omitempty"`
	Comment             string              `json:"comment,omitempty"`
	Due                 *string             `json:"due,omitempty"`
	EmergencyContact    EmergencyContact    `json:"emergency_contact,omitempty"`
	Items               Items               `json:"items"`
	OfferPayload        string              `json:"offer_payload,omitempty"`
	OptionalReturn      bool                `json:"optional_return,omitempty"`
	ReferralSource      string              `json:"referral_source,omitempty"`
	RoutePoints         RoutePointsCreate   `json:"route_points"`
	SameDayData         *SameDayData        `json:"same_day_data,omitempty"`
	ShippingDocument    string              `json:"shipping_document,omitempty"`
	SkipAct             bool                `json:"skip_act,omitempty"`
	SkipClientNotify    bool                `json:"skip_client_notify,omitempty"`
	SkipDoorToDoor      bool                `json:"skip_door_to_door,omitempty"`
	SkipEmergencyNotify bool                `json:"skip_emergency_notify,omitempty"`
}
type RoutePointsCreate []RoutePointCreate

type RoutePointCreate struct {
	Address           Address                   `json:"address"`
	Buyout            *create.Buyout            `json:"buyout,omitempty"`
	Contact           Contact                   `json:"contact"`
	ExternalOrderCost *ExternalOrderCost        `json:"external_order_cost,omitempty"`
	ExternalOrderID   string                    `json:"external_order_id,omitempty"`
	LeaveUnderDoor    bool                      `json:"leave_under_door,omitempty"`
	MeetOutside       bool                      `json:"meet_outside,omitempty"`
	NoDoorCall        bool                      `json:"no_door_call,omitempty"`
	PaymentOnDelivery *create.PaymentOnDelivery `json:"payment_on_delivery,omitempty"`
	PickupCode        string                    `json:"pickup_code,omitempty"`
	PointID           int64                     `json:"point_id,omitempty"`
	SkipConfirmation  bool                      `json:"skip_confirmation,omitempty"`
	Type              string                    `json:"type,omitempty"`
	VisitOrder        int64                     `json:"visit_order,omitempty"`
}
