package responses

import "github.com/vildan-valeev/yandexdelivery/models"

type APIResponseCourierPhone struct {
	models.CourierPhoneResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseCourierPhone) Base() APIResponseBase {
	return a.APIResponseBase
}
