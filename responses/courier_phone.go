package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseCourierPhone struct {
	models.CourierPhoneResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseCourierPhone) Base() APIResponseBase {
	return a.APIResponseBase
}
