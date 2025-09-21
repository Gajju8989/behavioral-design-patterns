package service

import "strategypattern/paymentprocess"

type PaymentService interface {
	Pay(method string, amount int) (string, error)
}

type impl struct {
	paymentProcess paymentprocess.ProcessPayment
}

func NewPaymentService(paymentProcess paymentprocess.ProcessPayment) PaymentService {
	return &impl{
		paymentProcess: paymentProcess,
	}
}
