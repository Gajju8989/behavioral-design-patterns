package service

import (
	"errors"
	"strategypattern/paymentstrategy"
)

var (
	errInvalidMethod = errors.New("Invalid method")
)

func (i *impl) Pay(method string, amount int) (string, error) {
	switch method {
	case "credit_card":
		i.paymentProcess.PaymentStrategy = &paymentstrategy.CreditCardDetail{}
	case "debit_card":
		i.paymentProcess.PaymentStrategy = &paymentstrategy.DebitCardDetail{}
	case "paypal":
		i.paymentProcess.PaymentStrategy = &paymentstrategy.PaypalAccountDetail{}
	default:
		return "", errInvalidMethod
	}

	return i.paymentProcess.PaymentStrategy.Pay(amount), nil
}
