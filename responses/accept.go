package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseAccept struct {
	models.AcceptResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseAccept) Base() APIResponseBase {
	return a.APIResponseBase
}
