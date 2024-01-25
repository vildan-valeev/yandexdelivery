package responses

import "github.com/vildan-valeev/yandexdelivery/models"

type APIResponseCheckPrice struct {
	models.CheckPriceResponse
	APIResponseBase
}

func (a APIResponseCheckPrice) Base() APIResponseBase {
	return a.APIResponseBase
}
