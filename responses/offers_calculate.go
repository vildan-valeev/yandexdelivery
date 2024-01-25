package responses

import "github.com/vildan-valeev/yandexdelivery/models"

type APIResponseOffersCalculate struct {
	models.OffersCalculateResponse
	APIResponseBase
}

func (a APIResponseOffersCalculate) Base() APIResponseBase {
	return a.APIResponseBase
}
