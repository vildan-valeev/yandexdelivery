package responses

import "github.com/vildan-valeev/yandexdelivery/models"

type APIResponseCourierPosition struct {
	models.CourierPositionResponse
	APIResponseBase
}

func (a APIResponseCourierPosition) Base() APIResponseBase {
	return a.APIResponseBase
}
