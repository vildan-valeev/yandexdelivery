package responses

import (
	"github.com/vildan-valeev/yandexdelivery/models"
)

type APIResponseTariffs struct {
	models.TariffsResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseTariffs) Base() APIResponseBase {
	return a.APIResponseBase
}
