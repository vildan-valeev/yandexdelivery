package yandexdelivery

import (
	m "github.com/vildan-valeev/yandexdelivery/models"
	o "github.com/vildan-valeev/yandexdelivery/options"
	r "github.com/vildan-valeev/yandexdelivery/responses"
)

type YandexDelivery interface {
	DeliveryMethods(token string, payload m.DeliveryMethodsRequest) (res r.APIResponseDeliveryMethods, err error)

	CheckPrice(token string, payload m.CheckPriceRequest) (res r.APIResponseCheckPrice, err error)
	Tariffs(token string, payload m.TariffsRequest) (res r.APIResponseTariffs, err error)
	OffersCalculate(token string, payload m.OffersCalculateRequest) (res r.APIResponseOffersCalculate, err error)

	Create(token string, payload m.CreateRequest, opts *o.CreateOptions) (res r.APIResponseInfo, err error)
	Accept(token string, payload m.AcceptRequest, opts *o.AcceptOptions) (res r.APIResponseAccept, err error)
	Search(token string, payload m.SearchRequest) (res r.APIResponseSearch, err error)
	Info(token string, payload m.InfoRequest, opts *o.InfoOptions) (res r.APIResponseInfo, err error)

	CancelInfo(token string, payload m.CancelInfoRequest, opts *o.CancelInfoOptions) (res r.APIResponseCancelInfo, err error)
	Cancel(token string, payload m.CancelRequest, opts *o.CancelOptions) (res r.APIResponseCancel, err error)
	Return(token string, payload m.ReturnRequest, opts *o.ReturnOptions) (res r.APIResponseReturn, err error)

	CourierPhone(token string, payload m.CourierPhoneRequest) (res r.APIResponseCourierPhone, err error)
	CourierPosition(token string, payload m.CourierPositionRequest, opts *o.CourierPositionOptions) (res r.APIResponseCourierPosition, err error)
	TrackingLinks(token string, payload m.TrackingLinksRequest, opts *o.TrackingLinksOptions) (res r.APIResponseTrackingLinks, err error)

	ConfirmationCode(token string, payload m.ConfirmationCodeRequest, opts *o.ConfirmationCodeOptions) (res r.APIResponseConfirmationCode, err error)
	Document(token string, payload m.DocumentRequest, opts *o.DocumentOptions) (res r.APIResponseDocument, err error)

	BulkInfo(token string, payload m.BulkInfoRequest) (res r.APIResponseBulkInfo, err error)
	Journal(token string, payload m.JournalRequest) (res r.APIResponseJournal, err error)
	PointsEta(token string, payload m.PointsEtaRequest, opts *o.PointsEtaOptions) (res r.APIResponsePointsEta, err error)

	Edit(token string, payload m.EditRequest, opts *o.EditOptions) (res r.APIResponseEdit, err error)
	ApplyChanges(token string, payload m.ApplyChangesRequest, opts *o.ApplyChangesOptions) (res r.APIResponseApplyChanges, err error)
	ApplyChangesResult(token string, payload m.ApplyChangesResultRequest, opts *o.ApplyChangesResultOptions) (res r.APIResponseApplyChangesResult, err error)

	RobotCheckAvailability(token string, payload m.RobotCheckAvailabilityRequest) (res r.APIResponseRobotCheckAvailability, err error)
	RobotOpen(token string, payload m.RobotOpenRequest, opts *o.RobotOpenOptions) (res r.APIResponseRobotOpen, err error)

	PhotosPoint(token string, payload m.PhotosPointRequest) (res r.APIResponsePhotosPoint, err error)
	ProofDelivery(token string, payload m.ProofDeliveryRequest, opts *o.ProofDeliveryOptions) (res r.APIResponseProofDelivery, err error)
}
