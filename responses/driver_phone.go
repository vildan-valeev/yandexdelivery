package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseDriverPhone struct {
	models.DriverPhoneResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseDriverPhone) Base() APIResponseBase {
	return a.APIResponseBase
}
