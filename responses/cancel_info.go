package responses

import "github.com/vildan-valeev/yandexdelivery/models"

type APIResponseCancelInfo struct {
	models.CancelInfoResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseCancelInfo) Base() APIResponseBase {
	return a.APIResponseBase
}
