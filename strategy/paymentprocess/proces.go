package paymentprocess

import "strategypattern/paymentstrategy"

type ProcessPayment struct {
	PaymentStrategy paymentstrategy.PaymentStrategy
}

func (s *ProcessPayment) SetPayment(paymentStrategy paymentstrategy.PaymentStrategy) {
	s.PaymentStrategy = paymentStrategy
}

func (s *ProcessPayment) Checkout(amount int) {
	s.PaymentStrategy.Pay(amount)
}
