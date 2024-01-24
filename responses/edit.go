package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseEdit struct {
	models.Claim
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseEdit) Base() APIResponseBase {
	return a.APIResponseBase
}
