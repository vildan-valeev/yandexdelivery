package yandexlogistic

import "github.com/vildan-valeev/yandexlogistic/models"

// APIResponse is implemented by all the APIResponse* types.
type APIResponse interface {
	// Base returns the object of type APIResponseBase contained in each implemented type.
	Base() APIResponseBase
}

type APIResponseBase struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

// Base returns the APIResponseBase itself.
func (a APIResponseBase) Base() APIResponseBase {
	return a
}

//type ResponseCreate struct {
//	ID                 string               `json:"id"`
//	RoutePoints        []models.RoutePoints `json:"route_points"`
//	Pricing            models.Pricing       `json:"pricing,omitempty"`
//	ClientRequirements Requirements         `json:"client_requirements,omitempty"`
//	Status             string               `json:"status"`
//	Version            int64                `json:"version"`
//}

//type APIResponseCreate struct {
//	InfoResponse
//	APIResponseBase
//}

//// Base returns the contained object of type APIResponseBase.
//func (a APIResponseCreate) Base() APIResponseBase {
//	return a.APIResponseBase
//}
//

type APIResponseAccept struct {
	models.ResponseAccept
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseAccept) Base() APIResponseBase {
	return a.APIResponseBase
}

type APIResponseInfo struct {
	models.InfoResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseInfo) Base() APIResponseBase {
	return a.APIResponseBase
}

type APIResponseCancel struct {
	models.ResponseCancel
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseCancel) Base() APIResponseBase {
	return a.APIResponseBase
}

//type CreateRequest struct {
//	RoutePoints        []models.RoutePoints `json:"route_points,omitempty"`
//	ClientRequirements Requirements         `json:"client_requirements,omitempty"`
//	Items              []models.Items       `json:"items,omitempty"`
//	EmergencyContact   Contact              `json:"emergency_contact,omitempty"`
//	Due                string               `json:"due,omitempty"`
//}

type AcceptRequest struct {
	Version int64 `json:"version"`
}

type CancelRequest struct {
	Version     int64  `json:"version,omitempty"`
	CancelState string `json:"cancel_state,omitempty"`
}

type StatusCode int
