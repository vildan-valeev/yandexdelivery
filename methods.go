package yandexlogistic

import (
	"encoding/json"
	"github.com/vildan-valeev/yandexlogistic/models"
	"github.com/vildan-valeev/yandexlogistic/options"
	"github.com/vildan-valeev/yandexlogistic/responses"
)

// DeliveryMethods 1.1. Интервалы «Доставки в течение дня» https://yandex.ru/dev/logistics/api/ref/same-day/IntegrationV2DeliveryMethods.html
func (yc *YandexClient) DeliveryMethods(token string, payload models.DeliveryMethodsRequest, opts *options.DeliveryMethodsOptions) (res responses.APIResponseDeliveryMethods, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseDeliveryMethods{}, err
	}

	return post[responses.APIResponseDeliveryMethods](token, yc.url, methodDeliveryMethods, addValues(nil, opts), jsn)
}

// CheckPrice 2.1. Первичная оценка стоимости без создания заявки https://yandex.ru/dev/logistics/api/ref/estimate/IntegrationV2CheckPrice.html
func (yc *YandexClient) CheckPrice(token string, payload models.CheckPriceRequest, opts *options.CheckPriceOptions) (res responses.APIResponseCheckPrice, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseCheckPrice{}, err
	}

	return post[responses.APIResponseCheckPrice](token, yc.url, methodCheckPrice, addValues(nil, opts), jsn)
}

// Create 3.1. Создание заявки. https://yandex.ru/dev/logistics/api/ref/basic/IntegrationV2ClaimsCreate.html
func (yc *YandexClient) Create(token string, payload models.CreateRequest, opts *options.CreateOptions) (res responses.APIResponseInfo, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseInfo{}, err
	}

	return post[responses.APIResponseInfo](token, yc.url, methodCreate, addValues(nil, opts), jsn)
}

// Accept 3.2. Подтверждение заявки https://yandex.ru/dev/logistics/api/ref/basic/IntegrationV2ClaimsAccept.html
func (yc *YandexClient) Accept(token string, payload models.AcceptRequest, opts *options.AcceptOptions) (res responses.APIResponseAccept, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseAccept{}, err
	}

	return post[responses.APIResponseAccept](token, yc.url, methodAccept, addValues(nil, opts), jsn)
}

func (yc *YandexClient) Info(token string, opts *options.InfoOptions) (res responses.APIResponseInfo, err error) {
	return post[responses.APIResponseInfo](token, yc.url, methodInfo, addValues(nil, opts), nil)
}

func (yc *YandexClient) Cancel(token string, payload models.CancelRequest, opts *options.CancelOptions) (res responses.APIResponseCancel, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseCancel{}, err
	}

	return post[responses.APIResponseCancel](token, yc.url, methodCancel, addValues(nil, opts), jsn)
}
