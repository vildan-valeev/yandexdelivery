package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseReturn struct {
	models.ReturnResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseReturn) Base() APIResponseBase {
	return a.APIResponseBase
}
