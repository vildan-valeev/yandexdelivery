package create

import "errors"

type Buyout struct {
	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`
}

type Customer struct {
	Email string `json:"email,omitempty"`
	Inn   string `json:"inn,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type PaymentOnDelivery struct {
	Customer      Customer      `json:"customer,omitempty"`
	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`
}

type PaymentMethod string

const (
	PaymentMethodCash = "cash"
	PaymentMethodCard = "card"
)

func ToPaymentMethod(i string) (PaymentMethod, error) {
	p := PaymentMethod(i)
	switch p {
	case PaymentMethodCash, PaymentMethodCard:
		return p, nil
	default:
		return "", errors.New("invalid payment method")
	}
}
