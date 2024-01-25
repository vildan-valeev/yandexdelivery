package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseApplyChanges struct {
	models.ApplyChangesResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseApplyChanges) Base() APIResponseBase {
	return a.APIResponseBase
}
