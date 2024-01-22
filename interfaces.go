package yandexlogistic

import "github.com/vildan-valeev/yandexlogistic/models"

type YandexLogistics interface {
	Create(token string, payload models.CreateRequest, opts *CreateOptions) (res APIResponseInfo, err error)
	Accept(token string, payload AcceptRequest, opts *AcceptOptions) (res APIResponseAccept, err error)
	Info(token string, opts *InfoOptions) (res APIResponseInfo, err error)
	Cancel(token string, payload CancelRequest, opts *CancelOptions) (res APIResponseCancel, err error)
}
