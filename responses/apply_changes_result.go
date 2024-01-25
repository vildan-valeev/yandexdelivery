package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseApplyChangesResult struct {
	models.ApplyChangesResultResponse
	APIResponseBase
}

func (a APIResponseApplyChangesResult) Base() APIResponseBase {
	return a.APIResponseBase
}
