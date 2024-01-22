package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseSearch struct {
	models.SearchResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseSearch) Base() APIResponseBase {
	return a.APIResponseBase
}
