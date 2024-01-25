package responses

import "github.com/vildan-valeev/yandexdelivery/models"

type APIResponseJournal struct {
	models.JournalResponse
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseJournal) Base() APIResponseBase {
	return a.APIResponseBase
}
