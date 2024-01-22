package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseOffersCalculate struct {
	models.OffersCalculateResponse
	APIResponseBase
}

func (a APIResponseOffersCalculate) Base() APIResponseBase {
	return a.APIResponseBase
}
