package responses

import "github.com/vildan-valeev/yandexdelivery/models"

type APIResponseBulkInfo struct {
	models.BulkInfoResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseBulkInfo) Base() APIResponseBase {
	return a.APIResponseBase
}
