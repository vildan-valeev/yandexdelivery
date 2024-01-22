package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseCourierPosition struct {
	models.CourierPositionResponse
	APIResponseBase
}

func (a APIResponseCourierPosition) Base() APIResponseBase {
	return a.APIResponseBase
}
