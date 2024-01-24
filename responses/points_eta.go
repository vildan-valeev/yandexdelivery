package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponsePointsEta struct {
	models.PointsEtaResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponsePointsEta) Base() APIResponseBase {
	return a.APIResponseBase
}
