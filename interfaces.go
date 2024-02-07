package yandexdelivery

import (
	"github.com/vildan-valeev/yandexdelivery/models"
	"github.com/vildan-valeev/yandexdelivery/options"
	"github.com/vildan-valeev/yandexdelivery/responses"
)

type YandexDelivery interface {
	DeliveryMethods(token string, payload models.DeliveryMethodsRequest) (responses.APIResponseDeliveryMethods, error)

	CheckPrice(token string, payload models.CheckPriceRequest) (responses.APIResponseCheckPrice, error)
	Tariffs(token string, payload models.TariffsRequest) (responses.APIResponseTariffs, error)
	OffersCalculate(token string, payload models.OffersCalculateRequest) (responses.APIResponseOffersCalculate, error)

	Create(token string, payload models.CreateRequest, opts *options.CreateOptions) (responses.APIResponseInfo, error)
	Accept(token string, payload models.AcceptRequest, opts *options.AcceptOptions) (responses.APIResponseAccept, error)
	Search(token string, payload models.SearchRequest) (responses.APIResponseSearch, error)
	Info(token string, opts *options.InfoOptions) (responses.APIResponseInfo, error)

	CancelInfo(token string, opts *options.CancelInfoOptions) (responses.APIResponseCancelInfo, error)
	Cancel(token string, payload models.CancelRequest, opts *options.CancelOptions) (responses.APIResponseCancel, error)
	Return(token string, payload models.ReturnRequest, opts *options.ReturnOptions) (responses.APIResponseReturn, error)

	CourierPhone(token string, payload models.CourierPhoneRequest) (responses.APIResponseCourierPhone, error)
	CourierPosition(token string, opts *options.CourierPositionOptions) (responses.APIResponseCourierPosition, error)
	TrackingLinks(token string, opts *options.TrackingLinksOptions) (responses.APIResponseTrackingLinks, error)

	ConfirmationCode(token string, payload models.ConfirmationCodeRequest) (responses.APIResponseConfirmationCode, error)
	Document(token string, opts *options.DocumentOptions) (responses.APIResponseDocument, error)

	BulkInfo(token string, payload models.BulkInfoRequest) (responses.APIResponseBulkInfo, error)
	Journal(token string, payload models.JournalRequest) (responses.APIResponseJournal, error)
	PointsEta(token string, opts *options.PointsEtaOptions) (responses.APIResponsePointsEta, error)

	Edit(token string, payload models.EditRequest, opts *options.EditOptions) (responses.APIResponseEdit, error)
	ApplyChanges(token string, payload models.ApplyChangesRequest, opts *options.ApplyChangesOptions) (responses.APIResponseApplyChanges, error)
	ApplyChangesResult(token string, payload models.ApplyChangesResultRequest, opts *options.ApplyChangesResultOptions) (responses.APIResponseApplyChangesResult, error)

	RobotCheckAvailability(token string, payload models.RobotCheckAvailabilityRequest) (responses.APIResponseRobotCheckAvailability, error)
	RobotOpen(token string, payload models.RobotOpenRequest, opts *options.RobotOpenOptions) (responses.APIResponseRobotOpen, error)

	PhotosPoint(token string, payload models.PhotosPointRequest) (responses.APIResponsePhotosPoint, error)
	ProofDelivery(token string, opts *options.ProofDeliveryOptions) (responses.APIResponseProofDelivery, error)
}
