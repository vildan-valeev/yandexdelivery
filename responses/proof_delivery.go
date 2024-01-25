package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseProofDelivery struct {
	models.ProofDeliveryResponse
	APIResponseBase
}

func (a APIResponseProofDelivery) Base() APIResponseBase {
	return a.APIResponseBase
}
