package yandexlogistic

import (
	"log"
	"net/url"
)

const (
	methodBase            = "b2b/cargo/integration/v2/"
	methodDeliveryMethods = "delivery-methods"
	methodCheckPrice      = "check-price"
	methodTariffs         = "tariffs"
	methodOffersCalculate = "offers/calculate"
	methodCreate          = "claims/create"
	methodAccept          = "claims/accept"
	methodSearch          = "claims/search"
	methodInfo            = "claims/info"
	methodCancelInfo      = "claims/cancel-info"
	methodCancel          = "claims/cancel"
	methodReturn          = "claims/return"
	methodCourierPhone    = "driver-voiceforwarding"
	methodCourierPosition = "claims/performer-position"
)

type YandexClient struct {
	url   string
	token string
}

func NewYandexClient(urlBase string) *YandexClient {
	// TODO: перенести токен в клиента
	u, err := url.JoinPath(urlBase, methodBase)
	if err != nil {
		log.Fatal(err)
	}
	return &YandexClient{
		url: u,
	}
}
