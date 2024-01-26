package yandexdelivery

import (
	"encoding/json"
	"github.com/vildan-valeev/yandexdelivery/models"
	"github.com/vildan-valeev/yandexdelivery/options"
	"github.com/vildan-valeev/yandexdelivery/responses"
)

// DeliveryMethods 1.1. Интервалы «Доставки в течение дня» https://yandex.ru/dev/logistics/api/ref/same-day/IntegrationV2DeliveryMethods.html
func (yc *YandexClient) DeliveryMethods(token string, payload models.DeliveryMethodsRequest) (res responses.APIResponseDeliveryMethods, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseDeliveryMethods{}, err
	}

	return post[responses.APIResponseDeliveryMethods](yc.debugMode, token, yc.url, methodDeliveryMethods, nil, jsn)
}

// CheckPrice 2.1. Первичная оценка стоимости без создания заявки https://yandex.ru/dev/logistics/api/ref/estimate/IntegrationV2CheckPrice.html
func (yc *YandexClient) CheckPrice(token string, payload models.CheckPriceRequest) (res responses.APIResponseCheckPrice, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseCheckPrice{}, err
	}

	return post[responses.APIResponseCheckPrice](yc.debugMode, token, yc.url, methodCheckPrice, nil, jsn)
}

// Tariffs 2.2. Получение тарифов, доступных в точке https://yandex.ru/dev/logistics/api/ref/estimate/IntegrationV2Tariffs.html
func (yc *YandexClient) Tariffs(token string, payload models.TariffsRequest) (res responses.APIResponseTariffs, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseTariffs{}, err
	}

	return post[responses.APIResponseTariffs](yc.debugMode, token, yc.url, methodTariffs, nil, jsn)
}

// OffersCalculate 2.3. Получение вариантов доставки https://yandex.ru/dev/logistics/api/ref/estimate/IntegrationV2OfferCalculate.html
func (yc *YandexClient) OffersCalculate(token string, payload models.OffersCalculateRequest) (res responses.APIResponseOffersCalculate, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseOffersCalculate{}, err
	}

	return post[responses.APIResponseOffersCalculate](yc.debugMode, token, yc.url, methodOffersCalculate, nil, jsn)
}

// Create 3.1. Создание заявки. https://yandex.ru/dev/logistics/api/ref/basic/IntegrationV2ClaimsCreate.html
func (yc *YandexClient) Create(token string, payload models.CreateRequest, opts *options.CreateOptions) (res responses.APIResponseInfo, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseInfo{}, err
	}

	return post[responses.APIResponseInfo](yc.debugMode, token, yc.url, methodCreate, addValues(nil, opts), jsn)
}

// Accept 3.2. Подтверждение заявки https://yandex.ru/dev/logistics/api/ref/basic/IntegrationV2ClaimsAccept.html
func (yc *YandexClient) Accept(token string, payload models.AcceptRequest, opts *options.AcceptOptions) (res responses.APIResponseAccept, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseAccept{}, err
	}

	return post[responses.APIResponseAccept](yc.debugMode, token, yc.url, methodAccept, addValues(nil, opts), jsn)
}

// Search 3.3. Поиск заявок https://yandex.ru/dev/logistics/api/ref/basic/IntegrationV2ClaimsSearch.html
func (yc *YandexClient) Search(token string, payload models.SearchRequest) (res responses.APIResponseSearch, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseSearch{}, err
	}

	return post[responses.APIResponseSearch](yc.debugMode, token, yc.url, methodSearch, nil, jsn)
}

// Info 3.4. Получение информации по заявке https://yandex.ru/dev/logistics/api/ref/basic/IntegrationV2ClaimsInfo.html
func (yc *YandexClient) Info(token string, opts *options.InfoOptions) (res responses.APIResponseInfo, err error) {
	return post[responses.APIResponseInfo](yc.debugMode, token, yc.url, methodInfo, addValues(nil, opts), nil)
}

// CancelInfo 4.1. Получение признака отмены https://yandex.ru/dev/logistics/api/ref/cancel-and-skip-points/IntegrationV2ClaimsCancelInfo.html
func (yc *YandexClient) CancelInfo(token string, opts *options.CancelInfoOptions) (res responses.APIResponseCancelInfo, err error) {
	return post[responses.APIResponseCancelInfo](yc.debugMode, token, yc.url, methodCancelInfo, addValues(nil, opts), nil)
}

// Cancel 4.2. Отмена заявки https://yandex.ru/dev/logistics/api/ref/cancel-and-skip-points/IntegrationV2ClaimsCancel.html
func (yc *YandexClient) Cancel(token string, payload models.CancelRequest, opts *options.CancelOptions) (res responses.APIResponseCancel, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseCancel{}, err
	}

	return post[responses.APIResponseCancel](yc.debugMode, token, yc.url, methodCancel, addValues(nil, opts), jsn)
}

// Return 4.3. Пропуск точки в заказе с мультиточками. https://yandex.ru/dev/logistics/api/ref/cancel-and-skip-points/IntegrationV2ClaimsReturn.html
func (yc *YandexClient) Return(token string, payload models.ReturnRequest, opts *options.ReturnOptions) (res responses.APIResponseReturn, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseReturn{}, err
	}

	return post[responses.APIResponseReturn](yc.debugMode, token, yc.url, methodReturn, addValues(nil, opts), jsn)
}

// CourierPhone 5.1. Получение номера телефона курьера https://yandex.ru/dev/logistics/api/ref/performer-info/IntegrationV2DriverVoiceForwarding.html
func (yc *YandexClient) CourierPhone(token string, payload models.CourierPhoneRequest) (res responses.APIResponseCourierPhone, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseCourierPhone{}, err
	}

	return post[responses.APIResponseCourierPhone](yc.debugMode, token, yc.url, methodCourierPhone, nil, jsn)
}

// CourierPosition 5.2. Получение местоположения курьера https://yandex.ru/dev/logistics/api/ref/performer-info/IntegrationV2ClaimsPerformerPosition.html
func (yc *YandexClient) CourierPosition(token string, opts *options.CourierPositionOptions) (res responses.APIResponseCourierPosition, err error) {
	return get[responses.APIResponseCourierPosition](yc.debugMode, token, yc.url, methodCourierPosition, addValues(nil, opts), nil)
}

// TrackingLinks 5.3. Получение ссылок для отслеживания курьера https://yandex.ru/dev/logistics/api/ref/performer-info/IntegrationV2ClaimsTrackingLinks.html
func (yc *YandexClient) TrackingLinks(token string, opts *options.TrackingLinksOptions) (res responses.APIResponseTrackingLinks, err error) {
	return get[responses.APIResponseTrackingLinks](yc.debugMode, token, yc.url, methodTrackingLinks, addValues(nil, opts), nil)
}

// ConfirmationCode 6.1. Получение кода подтверждения https://yandex.ru/dev/logistics/api/ref/confirmation-code-and-acts/IntegrationV2ClaimsConfirmationCode.html
func (yc *YandexClient) ConfirmationCode(token string, opts *options.ConfirmationCodeOptions) (res responses.APIResponseConfirmationCode, err error) {
	return post[responses.APIResponseConfirmationCode](yc.debugMode, token, yc.url, methodConfirmationCode, addValues(nil, opts), nil)
}

// Document 6.2. Получение акта приёма-передачи https://yandex.ru/dev/logistics/api/ref/confirmation-code-and-acts/IntegrationV2ClaimsDocument.html
func (yc *YandexClient) Document(token string, opts *options.DocumentOptions) (res responses.APIResponseDocument, err error) {
	// TODO: получение файла - переделать запрос!!!!!
	return getFile[responses.APIResponseDocument](yc.debugMode, token, yc.url, methodDocument, addValues(nil, opts), nil)
}

// BulkInfo 7.1. Получение информации по нескольким заявкам  https://yandex.ru/dev/logistics/api/ref/claim-info/IntegrationV2ClaimsBulkInfo.html
func (yc *YandexClient) BulkInfo(token string, payload models.BulkInfoRequest) (res responses.APIResponseBulkInfo, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseBulkInfo{}, err
	}

	return post[responses.APIResponseBulkInfo](yc.debugMode, token, yc.url, methodBulkInfo, nil, jsn)
}

// Journal 7.2. Журнал изменений заказов  https://yandex.ru/dev/logistics/api/ref/claim-info/IntegrationV2ClaimsJournal.html
func (yc *YandexClient) Journal(token string, payload models.JournalRequest) (res responses.APIResponseJournal, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseJournal{}, err
	}

	return post[responses.APIResponseJournal](yc.debugMode, token, yc.url, methodJournal, nil, jsn)
}

// PointsEta 7.3. Получение прогноза по времени прибытия на точки  https://yandex.ru/dev/logistics/api/ref/claim-info/IntegrationV2ClaimsPointsEta.html
func (yc *YandexClient) PointsEta(token string, opts *options.PointsEtaOptions) (res responses.APIResponsePointsEta, err error) {
	return post[responses.APIResponsePointsEta](yc.debugMode, token, yc.url, methodPointsEta, addValues(nil, opts), nil)
}

// Edit 8.1. Редактирование заявки до её подтверждения https://yandex.ru/dev/logistics/api/ref/claim-edit/IntegrationV2ClaimsEdit.html
func (yc *YandexClient) Edit(token string, payload models.EditRequest, opts *options.EditOptions) (res responses.APIResponseEdit, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseEdit{}, err
	}

	return post[responses.APIResponseEdit](yc.debugMode, token, yc.url, methodEdit, addValues(nil, opts), jsn)
}

// ApplyChanges 8.2. Частичное редактирование заявки после ее подтверждения https://yandex.ru/dev/logistics/api/ref/claim-edit/ClaimsApplyChangesRequest.html
func (yc *YandexClient) ApplyChanges(token string, payload models.ApplyChangesRequest, opts *options.ApplyChangesOptions) (res responses.APIResponseApplyChanges, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseApplyChanges{}, err
	}

	return post[responses.APIResponseApplyChanges](yc.debugMode, token, yc.url, methodApplyChanges, addValues(nil, opts), jsn)
}

// ApplyChangesResult 8.3. Получить результат применения изменений https://yandex.ru/dev/logistics/api/ref/claim-edit/ClaimsApplyChangesResult.html
func (yc *YandexClient) ApplyChangesResult(token string, payload models.ApplyChangesResultRequest, opts *options.ApplyChangesResultOptions) (res responses.APIResponseApplyChangesResult, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseApplyChangesResult{}, err
	}

	return post[responses.APIResponseApplyChangesResult](yc.debugMode, token, yc.url, methodApplyChangesResult, addValues(nil, opts), jsn)
}

// RobotCheckAvailability 9.1 Запрос на проверку возможности доставки ровером https://yandex.ru/dev/logistics/api/ref/claim-edit/ClaimsApplyChangesResult.html
func (yc *YandexClient) RobotCheckAvailability(token string, payload models.RobotCheckAvailabilityRequest) (res responses.APIResponseRobotCheckAvailability, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseRobotCheckAvailability{}, err
	}

	return post[responses.APIResponseRobotCheckAvailability](yc.debugMode, token, yc.url, methodRobotCheckAvailability, nil, jsn)
}

// RobotOpen 9.2 Запрос на открытие крышки ровера https://yandex.ru/dev/logistics/api/ref/robot/IntegrationV2ClaimsRobotOpenRequest.html
func (yc *YandexClient) RobotOpen(token string, payload models.RobotOpenRequest, opts *options.RobotOpenOptions) (res responses.APIResponseRobotOpen, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponseRobotOpen{}, err
	}

	return post[responses.APIResponseRobotOpen](yc.debugMode, token, yc.url, methodRobotOpen, addValues(nil, opts), jsn)
}

// PhotosPoint 10.1. Получение фотографий по точке https://yandex.ru/dev/logistics/api/ref/proof-of-delivery/IntegrationV2ClaimsPhotosByPoint.html
func (yc *YandexClient) PhotosPoint(token string, payload models.PhotosPointRequest) (res responses.APIResponsePhotosPoint, err error) {
	jsn, err := json.Marshal(payload)
	if err != nil {
		return responses.APIResponsePhotosPoint{}, err
	}

	return post[responses.APIResponsePhotosPoint](yc.debugMode, token, yc.url, methodPhotosPoint, nil, jsn)
}

// ProofDelivery 10.1. Получение фотографий по точке https://yandex.ru/dev/logistics/api/ref/proof-of-delivery/IntegrationV2ClaimsProofOfDeliveryInfo.html
func (yc *YandexClient) ProofDelivery(token string, opts *options.ProofDeliveryOptions) (res responses.APIResponseProofDelivery, err error) {
	return post[responses.APIResponseProofDelivery](yc.debugMode, token, yc.url, methodProofDelivery, addValues(nil, opts), nil)
}
