package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseCancel struct {
	models.CancelResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseCancel) Base() APIResponseBase {
	return a.APIResponseBase
}
