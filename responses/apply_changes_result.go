package responses

import "github.com/vildan-valeev/yandexdelivery/models"

type APIResponseApplyChangesResult struct {
	models.ApplyChangesResultResponse
	APIResponseBase
}

func (a APIResponseApplyChangesResult) Base() APIResponseBase {
	return a.APIResponseBase
}
