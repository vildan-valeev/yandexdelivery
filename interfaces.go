package yandexlogistic

import (
	"github.com/vildan-valeev/yandexlogistic/models"
	"github.com/vildan-valeev/yandexlogistic/options"
	"github.com/vildan-valeev/yandexlogistic/responses"
)

type YandexLogistics interface {
	Create(token string, payload models.CreateRequest, opts *options.CreateOptions) (res responses.APIResponseInfo, err error)
	Accept(token string, payload models.AcceptRequest, opts *options.AcceptOptions) (res responses.APIResponseAccept, err error)
	Info(token string, opts *options.InfoOptions) (res responses.APIResponseInfo, err error)
	Cancel(token string, payload models.CancelRequest, opts *options.CancelOptions) (res responses.APIResponseCancel, err error)
}
