package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseConfirmationCode struct {
	models.ConfirmationCodeResponse
	APIResponseBase
}

func (a APIResponseConfirmationCode) Base() APIResponseBase {
	return a.APIResponseBase
}
