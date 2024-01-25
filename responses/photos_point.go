package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponsePhotosPoint struct {
	models.PhotosPointResponse
	APIResponseBase
}

func (a APIResponsePhotosPoint) Base() APIResponseBase {
	return a.APIResponseBase
}
