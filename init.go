package yandexlogistic

import (
	"log"
	"net/url"
)

const (
	methodBase            = "b2b/cargo/integration/v2/"
	methodDeliveryMethods = "delivery-methods"
	methodCheckPrice      = "check-price"
	methodCreate          = "claims/create"
	methodAccept          = "claims/accept"
	methodInfo            = "claims/info"
	methodCancel          = "claims/cancel"
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
