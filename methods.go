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

// Tariffs 2.2. Получение тарифов, доступных в точке https://yandex.ru/dev/logistics/api/ref/estimate/IntegrationV2Tariffs.html
func (yc *YandexClient) Tariffs(token string, payload models.TariffsRequest, opts *options.TariffsOptions) (res responses.APIResponseTariffs, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseTariffs{}, err
	}

	return post[responses.APIResponseTariffs](token, yc.url, methodTariffs, addValues(nil, opts), jsn)
}

// OffersCalculate 2.3. Получение вариантов доставки https://yandex.ru/dev/logistics/api/ref/estimate/IntegrationV2OfferCalculate.html
func (yc *YandexClient) OffersCalculate(token string, payload models.OffersCalculateRequest, opts *options.OffersCalculateOptions) (res responses.APIResponseOffersCalculate, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseOffersCalculate{}, err
	}

	return post[responses.APIResponseOffersCalculate](token, yc.url, methodOffersCalculate, addValues(nil, opts), jsn)
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

// Search 3.3. Поиск заявок https://yandex.ru/dev/logistics/api/ref/basic/IntegrationV2ClaimsSearch.html
func (yc *YandexClient) Search(token string, payload models.SearchRequest, opts *options.SearchOptions) (res responses.APIResponseSearch, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseSearch{}, err
	}

	return post[responses.APIResponseSearch](token, yc.url, methodSearch, addValues(nil, opts), jsn)
}

// Info 3.4. Получение информации по заявке https://yandex.ru/dev/logistics/api/ref/basic/IntegrationV2ClaimsInfo.html
func (yc *YandexClient) Info(token string, payload models.InfoRequest, opts *options.InfoOptions) (res responses.APIResponseInfo, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseInfo{}, err
	}

	return post[responses.APIResponseInfo](token, yc.url, methodInfo, addValues(nil, opts), jsn)
}

// CancelInfo 4.1. Получение признака отмены https://yandex.ru/dev/logistics/api/ref/cancel-and-skip-points/IntegrationV2ClaimsCancelInfo.html
func (yc *YandexClient) CancelInfo(token string, payload models.CancelInfoRequest, opts *options.CancelInfoOptions) (res responses.APIResponseCancelInfo, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseCancelInfo{}, err
	}

	return post[responses.APIResponseCancelInfo](token, yc.url, methodCancelInfo, addValues(nil, opts), jsn)
}

// Cancel 4.2. Отмена заявки https://yandex.ru/dev/logistics/api/ref/cancel-and-skip-points/IntegrationV2ClaimsCancel.html
func (yc *YandexClient) Cancel(token string, payload models.CancelRequest, opts *options.CancelOptions) (res responses.APIResponseCancel, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseCancel{}, err
	}

	return post[responses.APIResponseCancel](token, yc.url, methodCancel, addValues(nil, opts), jsn)
}

// Return 4.3. Пропуск точки в заказе с мультиточками. https://yandex.ru/dev/logistics/api/ref/cancel-and-skip-points/IntegrationV2ClaimsReturn.html
func (yc *YandexClient) Return(token string, payload models.ReturnRequest, opts *options.ReturnOptions) (res responses.APIResponseReturn, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseReturn{}, err
	}

	return post[responses.APIResponseReturn](token, yc.url, methodReturn, addValues(nil, opts), jsn)
}

// CourierPhone 5.1. Получение номера телефона курьера https://yandex.ru/dev/logistics/api/ref/performer-info/IntegrationV2DriverVoiceForwarding.html
func (yc *YandexClient) CourierPhone(token string, payload models.CourierPhoneRequest, opts *options.CourierPhoneOptions) (res responses.APIResponseCourierPhone, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseCourierPhone{}, err
	}

	return post[responses.APIResponseCourierPhone](token, yc.url, methodCourierPhone, addValues(nil, opts), jsn)
}

// CourierPosition 5.2. Получение местоположения курьера https://yandex.ru/dev/logistics/api/ref/performer-info/IntegrationV2ClaimsPerformerPosition.html
func (yc *YandexClient) CourierPosition(token string, payload models.CourierPositionRequest, opts *options.CourierPositionOptions) (res responses.APIResponseCourierPosition, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseCourierPosition{}, err
	}

	return get[responses.APIResponseCourierPosition](token, yc.url, methodCourierPosition, addValues(nil, opts), jsn)
}

// TrackingLinks 5.3. Получение ссылок для отслеживания курьера https://yandex.ru/dev/logistics/api/ref/performer-info/IntegrationV2ClaimsTrackingLinks.html
func (yc *YandexClient) TrackingLinks(token string, payload models.TrackingLinksRequest, opts *options.TrackingLinksOptions) (res responses.APIResponseTrackingLinks, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseTrackingLinks{}, err
	}

	return get[responses.APIResponseTrackingLinks](token, yc.url, methodTrackingLinks, addValues(nil, opts), jsn)
}

// ConfirmationCode 6.1. Получение кода подтверждения https://yandex.ru/dev/logistics/api/ref/confirmation-code-and-acts/IntegrationV2ClaimsConfirmationCode.html
func (yc *YandexClient) ConfirmationCode(token string, payload models.ConfirmationCodeRequest, opts *options.ConfirmationCodeOptions) (res responses.APIResponseConfirmationCode, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseConfirmationCode{}, err
	}

	return post[responses.APIResponseConfirmationCode](token, yc.url, methodConfirmationCode, addValues(nil, opts), jsn)
}

// Document 6.2. Получение акта приёма-передачи https://yandex.ru/dev/logistics/api/ref/confirmation-code-and-acts/IntegrationV2ClaimsDocument.html
func (yc *YandexClient) Document(token string, payload models.DocumentRequest, opts *options.DocumentOptions) (res responses.APIResponseDocument, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseDocument{}, err
	}
	// TODO: получение файла - переделать запрос!!!!!
	return getFile[responses.APIResponseDocument](token, yc.url, methodDocument, addValues(nil, opts), jsn)
}

// BulkInfo 7.1. Получение информации по нескольким заявкам  https://yandex.ru/dev/logistics/api/ref/claim-info/IntegrationV2ClaimsBulkInfo.html
func (yc *YandexClient) BulkInfo(token string, payload models.BulkInfoRequest, opts *options.BulkInfoOptions) (res responses.APIResponseBulkInfo, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseBulkInfo{}, err
	}

	return post[responses.APIResponseBulkInfo](token, yc.url, methodBulkInfo, addValues(nil, opts), jsn)
}

// Journal 7.2. Журнал изменений заказов  https://yandex.ru/dev/logistics/api/ref/claim-info/IntegrationV2ClaimsJournal.html
func (yc *YandexClient) Journal(token string, payload models.JournalRequest, opts *options.JournalOptions) (res responses.APIResponseJournal, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseJournal{}, err
	}

	return post[responses.APIResponseJournal](token, yc.url, methodJournal, addValues(nil, opts), jsn)
}

// PointsEta 7.3. Получение прогноза по времени прибытия на точки  https://yandex.ru/dev/logistics/api/ref/claim-info/IntegrationV2ClaimsPointsEta.html
func (yc *YandexClient) PointsEta(token string, payload models.PointsEtaRequest, opts *options.PointsEtaOptions) (res responses.APIResponsePointsEta, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponsePointsEta{}, err
	}

	return post[responses.APIResponsePointsEta](token, yc.url, methodPointsEta, addValues(nil, opts), jsn)
}

// Edit 8.1. Редактирование заявки до её подтверждения https://yandex.ru/dev/logistics/api/ref/claim-edit/IntegrationV2ClaimsEdit.html
func (yc *YandexClient) Edit(token string, payload models.EditRequest, opts *options.EditOptions) (res responses.APIResponseEdit, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseEdit{}, err
	}

	return post[responses.APIResponseEdit](token, yc.url, methodEdit, addValues(nil, opts), jsn)
}

// ApplyChanges Edit 8.2. Частичное редактирование заявки после ее подтверждения https://yandex.ru/dev/logistics/api/ref/claim-edit/ClaimsApplyChangesRequest.html
func (yc *YandexClient) ApplyChanges(token string, payload models.ApplyChangesRequest, opts *options.ApplyChangesOptions) (res responses.APIResponseApplyChanges, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseApplyChanges{}, err
	}

	return post[responses.APIResponseApplyChanges](token, yc.url, methodApplyChanges, addValues(nil, opts), jsn)
}

// ApplyChangesResult Edit 8.3. Получить результат применения изменений https://yandex.ru/dev/logistics/api/ref/claim-edit/ClaimsApplyChangesResult.html
func (yc *YandexClient) ApplyChangesResult(token string, payload models.ApplyChangesResultRequest, opts *options.ApplyChangesResultOptions) (res responses.APIResponseApplyChangesResult, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseApplyChangesResult{}, err
	}

	return post[responses.APIResponseApplyChangesResult](token, yc.url, methodApplyChangesResult, addValues(nil, opts), jsn)
}

// RobotCheckAvailability Edit 9.1 Запрос на проверку возможности доставки ровером https://yandex.ru/dev/logistics/api/ref/claim-edit/ClaimsApplyChangesResult.html
func (yc *YandexClient) RobotCheckAvailability(token string, payload models.RobotCheckAvailabilityRequest, opts *options.RobotCheckAvailabilityOptions) (res responses.APIResponseRobotCheckAvailability, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseRobotCheckAvailability{}, err
	}

	return post[responses.APIResponseRobotCheckAvailability](token, yc.url, methodRobotCheckAvailability, addValues(nil, opts), jsn)
}

// RobotOpen Edit 9.2 Запрос на открытие крышки ровера https://yandex.ru/dev/logistics/api/ref/robot/IntegrationV2ClaimsRobotOpenRequest.html
func (yc *YandexClient) RobotOpen(token string, payload models.RobotOpenRequest, opts *options.RobotOpenOptions) (res responses.APIResponseRobotOpen, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseRobotOpen{}, err
	}

	return post[responses.APIResponseRobotOpen](token, yc.url, methodRobotOpen, addValues(nil, opts), jsn)
}
