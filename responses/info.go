package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseInfo struct {
	models.Claim
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseInfo) Base() APIResponseBase {
	return a.APIResponseBase
}
