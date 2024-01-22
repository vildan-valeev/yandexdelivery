package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseDeliveryMethods struct {
	models.AcceptResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseDeliveryMethods) Base() APIResponseBase {
	return a.APIResponseBase
}
