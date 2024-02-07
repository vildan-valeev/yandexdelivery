package models

type CancelInfoResponse struct {
	CancelState  string `json:"cancel_state"`
	Currency     string `json:"currency"`
	Price        string `json:"price"`
	PriceWithVat string `json:"price_with_vat"`
}
