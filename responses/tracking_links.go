package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseTrackingLinks struct {
	models.TrackingLinksResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseTrackingLinks) Base() APIResponseBase {
	return a.APIResponseBase
}
