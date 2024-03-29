package yandexdelivery

import (
	"log"
	"net/url"
)

const (
	methodBase = "b2b/cargo/integration/v2/"

	methodDeliveryMethods        = "delivery-methods"
	methodCheckPrice             = "check-price"
	methodTariffs                = "tariffs"
	methodOffersCalculate        = "offers/calculate"
	methodCreate                 = "claims/create"
	methodAccept                 = "claims/accept"
	methodSearch                 = "claims/search"
	methodInfo                   = "claims/info"
	methodCancelInfo             = "claims/cancel-info"
	methodCancel                 = "claims/cancel"
	methodReturn                 = "claims/return"
	methodCourierPhone           = "driver-voiceforwarding"
	methodCourierPosition        = "claims/performer-position"
	methodTrackingLinks          = "claims/tracking-links"
	methodConfirmationCode       = "claims/confirmation_code"
	methodDocument               = "claims/document"
	methodBulkInfo               = "claims/bulk_info"
	methodJournal                = "claims/journal"
	methodPointsEta              = "claims/points-eta"
	methodEdit                   = "claims/edit"
	methodApplyChanges           = "claims/apply-changes/request"
	methodApplyChangesResult     = "claims/apply-changes/result"
	methodRobotCheckAvailability = "claims/robot/check-availability"
	methodRobotOpen              = "claims/robot/open-request"
	methodPhotosPoint            = "claims/photos_by_point"
	methodProofDelivery          = "claims/proof-of-delivery/info"
)

type YandexClient struct {
	debugMode bool
	url       string
	token     string
}

func NewYandexClient(urlBase, token string, debugMode bool) *YandexClient {
	u, err := url.JoinPath(urlBase, methodBase)
	if err != nil {
		log.Fatal(err)
	}
	return &YandexClient{
		url:       u,
		debugMode: debugMode,
		token:     token,
	}
}
