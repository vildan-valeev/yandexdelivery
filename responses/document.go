package responses

import "github.com/vildan-valeev/yandexdelivery/models"

type APIResponseDocument struct {
	models.DocumentResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseDocument) Base() APIResponseBase {
	return a.APIResponseBase
}
