package yandexdelivery

import (
	"github.com/vildan-valeev/yandexdelivery/models"
	"github.com/vildan-valeev/yandexdelivery/options"
	"github.com/vildan-valeev/yandexdelivery/responses"
)

type YandexDelivery interface {
	DeliveryMethods(token string, payload models.DeliveryMethodsRequest) (res responses.APIResponseDeliveryMethods, err error)

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

	CourierPhone(token string, payload models.CourierPhoneRequest, opts *options.CourierPhoneOptions) (res responses.APIResponseCourierPhone, err error)
	CourierPosition(token string, payload models.CourierPositionRequest, opts *options.CourierPositionOptions) (res responses.APIResponseCourierPosition, err error)
	TrackingLinks(token string, payload models.TrackingLinksRequest, opts *options.TrackingLinksOptions) (res responses.APIResponseTrackingLinks, err error)
}
