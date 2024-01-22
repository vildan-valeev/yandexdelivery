package yandexlogistic

import (
	"log"
	"net/url"
)

const (
	methodBase   = "b2b/cargo/integration/v2/claims/"
	methodCreate = "create"
	methodAccept = "accept"
	methodInfo   = "info"
	methodCancel = "cancel"
)

type YandexClient struct {
	url   string
	token string
}

func NewYandexClient(urlBase string) *YandexClient {
	u, err := url.JoinPath(urlBase, methodBase)
	if err != nil {
		log.Fatal(err)
	}
	return &YandexClient{
		url: u,
	}
}
