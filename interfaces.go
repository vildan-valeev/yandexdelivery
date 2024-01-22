package yandexlogistic

import (
	"github.com/vildan-valeev/yandexlogistic/models"
	"github.com/vildan-valeev/yandexlogistic/options"
	"github.com/vildan-valeev/yandexlogistic/responses"
)

type YandexLogistics interface {
	DeliveryMethods(token string, payload models.DeliveryMethodsRequest, opts *options.DeliveryMethodsOptions) (res responses.APIResponseDeliveryMethods, err error)

	CheckPrice(token string, payload models.CheckPriceRequest, opts *options.CheckPriceOptions) (res responses.APIResponseCheckPrice, err error)
	Tariffs(token string, payload models.TariffsRequest, opts *options.TariffsOptions) (res responses.APIResponseTariffs, err error)
	OffersCalculate(token string, payload models.OffersCalculateRequest, opts *options.OffersCalculateOptions) (res responses.APIResponseOffersCalculate, err error)

	Create(token string, payload models.CreateRequest, opts *options.CreateOptions) (res responses.APIResponseInfo, err error)
	Accept(token string, payload models.AcceptRequest, opts *options.AcceptOptions) (res responses.APIResponseAccept, err error)
	Search(token string, payload models.SearchRequest, opts *options.SearchOptions) (res responses.APIResponseSearch, err error)
	Info(token string, payload models.InfoRequest, opts *options.InfoOptions) (res responses.APIResponseInfo, err error)

	CancelInfo(token string, payload models.CancelInfoRequest, opts *options.CancelInfoOptions) (res responses.APIResponseCancelInfo, err error)
	Cancel(token string, payload models.CancelRequest, opts *options.CancelOptions) (res responses.APIResponseCancel, err error)
	Return(token string, payload models.ReturnRequest, opts *options.ReturnOptions) (res responses.APIResponseReturn, err error)

	DriverPhone(token string, payload models.DriverPhoneRequest, opts *options.DriverPhoneOptions) (res responses.APIResponseDriverPhone, err error)
}
